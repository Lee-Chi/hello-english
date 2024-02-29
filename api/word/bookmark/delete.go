package bookmark

import (
	"net/http"

	"hello-english/base/api"
	"hello-english/db"
	"hello-english/db/model"

	"github.com/Lee-Chi/go-sdk/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (g Group) Delete(ctx *gin.Context) {
	var request struct {
		UserID string `json:"userId"`
		WordID string `json:"wordId"`
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

	wordID, err := primitive.ObjectIDFromHex(request.WordID)
	if err != nil {
		code := api.ArgumentError
		logger.Error(code.Dump("Invalid request, parse feild wordId, %v", err))
		ctx.JSON(http.StatusBadRequest, code.Response())
		return
	}

	collection := db.Collection(model.CWordBookmark).
		Where(model.Key.WordBookmark.UserID.Eq(userID).And(model.Key.WordBookmark.ID.Eq(wordID)))

	if err := collection.DeleteOne(ctx); err != nil {
		code := api.DatabaseError
		logger.Error(code.Dump("Delete word bookmark error: %v", err))
		ctx.JSON(http.StatusInternalServerError, code.Response())
		return
	}

	ctx.JSON(http.StatusOK, response)
	return
}
