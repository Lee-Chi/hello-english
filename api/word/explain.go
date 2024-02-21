package word

import (
	"encoding/json"
	"fmt"
	"hello-english/base/api"
	"hello-english/base/openai"
	"net/http"
	"strings"
	"sync"

	"github.com/Lee-Chi/go-sdk/logger"
	"github.com/gin-gonic/gin"
)

var once sync.Once
var template string

func (g Group) Explain(ctx *gin.Context) {
	type Word struct {
		PartOfSpeech string   `json:"partOfSpeech"`
		Translation  string   `json:"translation"`
		Sentences    []string `json:"sentences"`
	}

	var request struct {
		Word string `json:"word" binding:"required"`
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

	once.Do(func() {
		example := []Word{
			{
				PartOfSpeech: "vt",
				Translation:  "to shout or speak loudly and insistently",
				Sentences: []string{
					"The dog barked at the intruder.", "The coach barked orders at the players.", "He barked out a command to stop.",
				},
			},
		}

		t, _ := json.Marshal(example)
		template = string(t)
	})

	content := fmt.Sprintf(`Translate the following words and provide port of speech and three example. output example: %s, input: %s`, template, request.Word)

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
