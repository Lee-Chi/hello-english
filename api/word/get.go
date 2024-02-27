package word

import (
	"net/http"

	"hello-english/base/api"
	"hello-english/db"
	"hello-english/db/model"

	"github.com/Lee-Chi/go-sdk/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (g Group) Get(ctx *gin.Context) {
	var request struct {
		Type string
	}

	var response struct {
		api.ResponseBase
		Word Word `json:"word"`
	}

	request.Type = ctx.Request.URL.Query().Get("type")

	collection := db.Collection(model.CWordBase).Sort(model.Key.WordBase.ID.Asc())

	if request.Type == model.Type_Sight {
		var wordSight struct {
			ID              primitive.ObjectID `bson:"_id"`
			model.WordSight `bson:"-,inline"`
		}
		if err := db.Collection(model.CWordSight).Sort(model.Key.WordSight.ID.Asc()).FindOne(ctx, &wordSight); err != nil {
			code := api.DatabaseError
			logger.Error(code.Dump("Find word sight error: %v", err))
			ctx.JSON(http.StatusInternalServerError, code.Response())
			return
		}

		collection = collection.Where(model.Key.WordSight.Letters.Eq(wordSight.Letters))
		// if sight word, return sight word id
		response.Word.ID = wordSight.ID.Hex()
	}

	var word struct {
		ID         primitive.ObjectID `bson:"_id"`
		model.Word `bson:"-,inline"`
	}
	if err := collection.FindOne(ctx, &word); err != nil {
		code := api.DatabaseError
		logger.Error(code.Dump("Find word error: %v", err))
		ctx.JSON(http.StatusInternalServerError, code.Response())
		return
	}

	if response.Word.ID == "" {
		// if not sight word, return word id
		response.Word.ID = word.ID.Hex()
	}

	response.Word.Letters = word.Letters
	for _, definition := range word.Definitions {
		response.Word.Definitions = append(response.Word.Definitions, Definition{
			PartOfSpeech: definition.PartOfSpeech,
			Translation:  definition.Translation,
		})
	}

	ctx.JSON(http.StatusOK, response)
	return
}
