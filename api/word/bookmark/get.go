package bookmark

import (
	"net/http"
	"strconv"

	"hello-english/base/api"
	"hello-english/db"
	"hello-english/db/model"

	"github.com/Lee-Chi/go-sdk/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (g Group) Get(ctx *gin.Context) {
	type WordBookmark struct {
		ID string `json:"id"`
		Word
	}

	var request struct {
		UserID   primitive.ObjectID `json:"userId"`
		Page     int64              `json:"page"`
		PageSize int64              `json:"pageSize"`
	}

	var response struct {
		api.ResponseBase
		Count int64          `json:"count"`
		Words []WordBookmark `json:"words"`
	}

	page, err := strconv.ParseInt(ctx.Request.URL.Query().Get("page"), 10, 64)
	if err != nil {
		code := api.ArgumentError
		logger.Error(code.Dump("Invalid request, parse feild page, %v", err))
		ctx.JSON(http.StatusBadRequest, code.Response())
		return
	}

	pageSize, err := strconv.ParseInt(ctx.Request.URL.Query().Get("pageSize"), 10, 64)
	if err != nil {
		code := api.ArgumentError
		logger.Error(code.Dump("Invalid request, parse feild pageSize, %v", err))
		ctx.JSON(http.StatusBadRequest, code.Response())
		return
	}

	userID, err := primitive.ObjectIDFromHex(ctx.Request.URL.Query().Get("userId"))
	if err != nil {
		code := api.ArgumentError
		logger.Error(code.Dump("Invalid request, parse feild userId, %v", err))
		ctx.JSON(http.StatusBadRequest, code.Response())
		return
	}

	request.UserID = userID
	request.Page = page
	request.PageSize = pageSize

	collection := db.Collection(model.CWordBookmark).Where(model.Key.WordBookmark.UserID.Eq(request.UserID))

	count, err := collection.Count(ctx)
	if err != nil {
		code := api.DatabaseError
		logger.Error(code.Dump("Count word bookmark error: %v", err))
		ctx.JSON(http.StatusInternalServerError, code.Response())
		return
	}
	response.Count = count

	skip := (request.Page - 1) * request.PageSize
	limit := request.PageSize

	collection = collection.Sort(model.Key.WordBookmark.ID.Desc()).Skip(skip).Limit(limit)

	var wordBookmarks []struct {
		ID                 primitive.ObjectID `bson:"_id"`
		model.WordBookmark `bson:"-,inline"`
	}
	if err := collection.Find(ctx, &wordBookmarks); err != nil {
		code := api.DatabaseError
		logger.Error(code.Dump("Find word bookmark error: %v", err))
		ctx.JSON(http.StatusInternalServerError, code.Response())
		return
	}

	for _, wordBookmark := range wordBookmarks {
		response.Words = append(response.Words, WordBookmark{
			ID: wordBookmark.ID.Hex(),
			Word: Word{
				Letters:      wordBookmark.Letters,
				PartOfSpeech: wordBookmark.PartOfSpeech,
				Translation:  wordBookmark.Translation,
				Sentences:    wordBookmark.Sentences,
			},
		})
	}

	ctx.JSON(http.StatusOK, response)
	return
}
