package sentence

import (
	"fmt"
	"hello-english/base/api"
	"hello-english/base/openai"
	"net/http"

	"github.com/Lee-Chi/go-sdk/logger"
	"github.com/gin-gonic/gin"
)

func (g Group) Translate(ctx *gin.Context) {
	var request struct {
		Sentence string `json:"sentence" binding:"required"`
	}

	var response struct {
		api.ResponseBase
		Translation string `json:"translation"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		code := api.ArgumentError
		logger.Error(code.Dump("Invalid request: %v", err))
		ctx.JSON(http.StatusBadRequest, code.Response())
		return
	}

	content := fmt.Sprintf(`翻譯以下英文句子為繁體中文，如果不是句子，回答'無法翻譯': %s`, request.Sentence)

	reply, err := openai.Chat(ctx, []openai.ChatCompletionMessage{
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

	response.Translation = reply
	ctx.JSON(http.StatusOK, response)
	return
}
