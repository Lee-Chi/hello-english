package bookmark

import (
	"net/http"
	"time"

	"hello-english/base/api"
	"hello-english/db"
	"hello-english/db/model"

	"github.com/Lee-Chi/go-sdk/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (g Group) Add(ctx *gin.Context) {
	var request struct {
		UserID string `json:"userId"`
		Word   Word   `json:"word"`
	}

	var response struct {
		api.ResponseBase
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		code := api.ArgumentError
		logger.Error(code.Dump("Invalid request: %v", err))
		ctx.JSON(http.StatusBadRequest, code.Response())
		return
	}

	userID, err := primitive.ObjectIDFromHex(request.UserID)
	if err != nil {
		code := api.ArgumentError
		logger.Error(code.Dump("Invalid request, parse feild userId, %v", err))
		ctx.JSON(http.StatusBadRequest, code.Response())
		return
	}

	collection := db.Collection(model.CWordBookmark).Where(model.Key.WordBookmark.UserID.Eq(userID))

	count, err := collection.Count(ctx)
	if err != nil {
		code := api.DatabaseError
		logger.Error(code.Dump("Count word bookmark error: %v", err))
		ctx.JSON(http.StatusInternalServerError, code.Response())
		return
	}

	wordBookmark := model.WordBookmark{
		UserID:    userID,
		Index:     count + 1,
		CreatedAt: time.Now(),
		WordExplained: model.WordExplained{
			Letters:      request.Word.Letters,
			PartOfSpeech: request.Word.PartOfSpeech,
			Translation:  request.Word.Translation,
			Sentences:    request.Word.Sentences,
		},
	}

	if _, err := collection.InsertOne(ctx, wordBookmark); err != nil {
		code := api.DatabaseError
		logger.Error(code.Dump("Insert word bookmark error: %v", err))
		ctx.JSON(http.StatusInternalServerError, code.Response())
		return
	}

	ctx.JSON(http.StatusOK, response)
	return
}
