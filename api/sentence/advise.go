package sentence

import (
	"encoding/json"
	"fmt"
	"hello-english/base/api"
	"hello-english/base/openai"
	"net/http"

	"github.com/Lee-Chi/go-sdk/logger"
	"github.com/gin-gonic/gin"
)

const Template string = `{"original":"Can you teach me English?","revised":"Can you teach me English?","reasons":[{"type":"grammar","message":"The sentence is grammatically correct."},{"type":"style","message":"The sentence is clear and concise."},{"type":"tone","message":"The sentence is polite and respectful."}]}`

func (g Group) Advise(ctx *gin.Context) {
	type Advice struct {
		Original string `json:"original"`
		Revised  string `json:"revised"`
		Reasons  []struct {
			Type    string `json:"type"`
			Message string `json:"message"`
		} `json:"reasons"`
	}

	var request struct {
		Sentence string `json:"sentence" binding:"required"`
	}

	var response struct {
		api.ResponseBase
		Advice Advice `json:"advice"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		code := api.ArgumentError
		logger.Error(code.Dump("Invalid request: %v", err))
		ctx.JSON(http.StatusBadRequest, code.Response())
		return
	}

	content := fmt.Sprintf(`Check the following sentences, output example: %s。 input: %s`, Template, request.Sentence)

	reply, err := openai.Chat(ctx, []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are an English teacher designed to output JSON.",
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

	if err := json.Unmarshal([]byte(reply), &response.Advice); err != nil {
		code := api.UnknownError
		logger.Error(code.Dump("Unmarshal reply: %s, error: %v", reply, err))
		ctx.JSON(http.StatusInternalServerError, code.Response())
		return
	}

	ctx.JSON(http.StatusOK, response)
	return
}
