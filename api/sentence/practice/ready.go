package practice

import (
	"encoding/json"
	"fmt"
	"hello-english/base/api"
	"hello-english/base/openai"
	"net/http"

	"github.com/Lee-Chi/go-sdk/logger"
	"github.com/gin-gonic/gin"
)

func (g Group) Ready(ctx *gin.Context) {
	type Practice struct {
		Chinese string `json:"chinese"`
	}

	var request struct {
		Topic string `json:"topic" binding:"required"`
	}

	var response struct {
		api.ResponseBase
		Question string `json:"question"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		code := api.ArgumentError
		logger.Error(code.Dump("Invalid request: %v", err))
		ctx.JSON(http.StatusBadRequest, code.Response())
		return
	}

	content := fmt.Sprintf(`Give me a random Traditional Chinese sentence about %s from wiki. output example:{"chinese": "這個條件式成立嗎？"} `, request.Topic)

	reply, err := openai.Chat(ctx, []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are an English teacher designed to output JSON. Now you are going to do a Traditional Chinese to English translation exercise.",
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

	var practice Practice
	if err := json.Unmarshal([]byte(reply), &practice); err != nil {
		code := api.UnknownError
		logger.Error(code.Dump("Unmarshal error: %v", err))
		ctx.JSON(http.StatusInternalServerError, code.Response())
		return
	}

	response.Question = practice.Chinese

	ctx.JSON(http.StatusOK, response)
	return
}
