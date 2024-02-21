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

const Template string = `{"correct_translation": "How do I invoke this function?", "advices":["The word 'call' can be replaced with 'invoke' or 'execute'.","The word 'function' can be replaced with 'method' or 'procedure'."]}`

func (g Group) Submit(ctx *gin.Context) {
	type Practice struct {
		CorrectTranslation string   `json:"correct_translation"`
		Advices            []string `json:"advices"`
	}

	var request struct {
		Question string `json:"question" binding:"required"`
		Answer   string `json:"answer" binding:"required"`
	}

	var response struct {
		api.ResponseBase
		CorrectAnswer string   `json:"correctAnswer"`
		Advices       []string `json:"advices"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		code := api.ArgumentError
		logger.Error(code.Dump("Invalid request: %v", err))
		ctx.JSON(http.StatusBadRequest, code.Response())
		return
	}

	content := fmt.Sprintf(`Is the translation of %s is %s, provide the correct tranlation and three advices. output example:%s`, request.Question, request.Answer, Template)

	reply, err := openai.Chat(ctx, []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are an English teacher designed to output JSON. Now you are going to do a Trandition Chinese to English translation exercise.",
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

	response.CorrectAnswer = practice.CorrectTranslation
	response.Advices = practice.Advices

	ctx.JSON(http.StatusOK, response)
	return
}
