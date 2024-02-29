package word

import (
	"encoding/json"
	"fmt"
	"hello-english/base/api"
	"hello-english/base/openai"
	"net/http"
	"strings"

	"github.com/Lee-Chi/go-sdk/logger"
	"github.com/gin-gonic/gin"
)

var example string = `[{"letters":"overwork ","partOfSpeech":"v","translation":"過度勞累","reason":"這個單字的意思是「工作過度」，可以用來描述工作很久都沒有休息的狀態。"},{"letters":"work without rest","partOfSpeech":"phrase","translation":"不休息地工作","reason":"這個短語的意思是「不休息地工作」，可以用來描述工作很久都沒有休息的具體情況。"},{"letters":"tireless ","partOfSpeech":"adj","translation":"不知疲倦","reason":"這個單字的意思是「不知疲倦」，可以用來描述工作很久都沒有休息的人的狀態。"}]`

func (g Group) Search(ctx *gin.Context) {
	type Word struct {
		Letters      string `json:"letters"`
		PartOfSpeech string `json:"partOfSpeech"`
		Translation  string `json:"translation"`
		Reason       string `json:"reason"`
	}

	var request struct {
		Prompt string `json:"prompt" binding:"required"`
	}

	var response struct {
		api.ResponseBase
		Words []Word `json:"words"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		code := api.ArgumentError
		logger.Error(code.Dump("Invalid request: %v", err))
		ctx.JSON(http.StatusBadRequest, code.Response())
		return
	}

	content := fmt.Sprintf(`Summarize the following text, and then provide three suitable words and give reasons for the recommendations. output example: %s, input: %s`, example, request.Prompt)

	reply, err := openai.Chat(ctx, []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are an English teacher designed to output JSON array.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		},
	})
	if err != nil {
		code := api.UnknownError
		logger.Error(code.Dump("Chat error: %v", err))
		ctx.JSON(http.StatusInternalServerError, code.Response())
		return
	}

	fmt.Println("reply:", reply)

	if strings.HasPrefix(reply, "```json") {
		reply = strings.TrimPrefix(reply, "```json")
		reply = strings.TrimSuffix(reply, "```")
	}

	words := make([]Word, 0)
	if err := json.Unmarshal([]byte(reply), &words); err != nil {
		code := api.UnknownError
		logger.Error(code.Dump("Unmarshal reply(%s) error: %v", reply, err))
		ctx.JSON(http.StatusInternalServerError, code.Response())
		return
	}

	response.Words = words
	ctx.JSON(http.StatusOK, response)
	return
}
