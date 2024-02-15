package word

import (
	"net/http"

	"hello-english/base/api"
	"hello-english/base/english"
	"hello-english/db"
	"hello-english/db/model"

	"github.com/Lee-Chi/go-sdk/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (g Group) Backward(ctx *gin.Context) {
	var request struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	}

	var response struct {
		api.ResponseBase
		Word Word `json:"word"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		code := api.ArgumentError
		logger.Error(code.Dump("Invalid request: %v", err))
		ctx.JSON(http.StatusBadRequest, code.Response())
		return
	}

	id, _ := primitive.ObjectIDFromHex(request.ID)

	colWord := db.Collection(model.CWord)

	switch request.Type {
	case english.Type_Sight:
		colWordSight := db.Collection(model.CWordSight).Sort(english.Word_Key_ID.Desc()).Where(english.Word_Key_ID.Lt(id))
		count, err := colWordSight.Count(ctx)
		if err != nil {
			code := api.DatabaseError
			logger.Error(code.Dump("Count word sight error: %v", err))
			ctx.JSON(http.StatusInternalServerError, code.Response())
			return
		}

		if count == 0 {
			// if not found, use the last sight word
			colWordSight = colWordSight.Where()
		}

		var wordSight struct {
			ID              primitive.ObjectID `bson:"_id"`
			model.WordSight `bson:"-,inline"`
		}

		if err := colWordSight.FindOne(ctx, &wordSight); err != nil {
			code := api.DatabaseError
			logger.Error(code.Dump("Find word sight error: %v", err))
			ctx.JSON(http.StatusInternalServerError, code.Response())
			return
		}

		letters := wordSight.Letters

		colWord = colWord.Where(english.Word_Key_Letters.Eq(letters))

		var word struct {
			ID         primitive.ObjectID `bson:"_id"`
			model.Word `bson:"-,inline"`
		}
		if err := colWord.FindOne(ctx, &word); err != nil {
			code := api.DatabaseError
			logger.Error(code.Dump("Find word error: %v", err))
			ctx.JSON(http.StatusInternalServerError, code.Response())
			return
		}

		response.Word.Letters = word.Letters
		for _, definition := range word.Definitions {
			response.Word.Definitions = append(response.Word.Definitions, Definition{
				PartOfSpeech: definition.PartOfSpeech,
				Translation:  definition.Translation,
			})
		}

		response.Word.ID = wordSight.ID.Hex()
	default:
		// if not sight word, use the previous word
		colWord = colWord.Where(english.Word_Key_ID.Lt(id)).Sort(english.Word_Key_ID.Desc())

		var word struct {
			ID         primitive.ObjectID `bson:"_id"`
			model.Word `bson:"-,inline"`
		}
		if err := colWord.FindOne(ctx, &word); err != nil {
			code := api.DatabaseError
			logger.Error(code.Dump("Find word error: %v", err))
			ctx.JSON(http.StatusInternalServerError, code.Response())
			return
		}

		response.Word.Letters = word.Letters
		for _, definition := range word.Definitions {
			response.Word.Definitions = append(response.Word.Definitions, Definition{
				PartOfSpeech: definition.PartOfSpeech,
				Translation:  definition.Translation,
			})
		}

		response.Word.ID = word.ID.Hex()
	}

	ctx.JSON(http.StatusOK, response)
	return
}
