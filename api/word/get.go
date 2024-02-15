package word

import (
	"net/http"
	"time"

	"hello-english/base/api"
	"hello-english/base/english"
	"hello-english/db"
	"hello-english/db/model"

	"github.com/Lee-Chi/go-sdk/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Definition struct {
	PartOfSpeech string `json:"partOfSpeech"`
	Translation  string `json:"translation"`
}

type Word struct {
	ID          string       `json:"id"`
	Letters     string       `json:"letters"`
	Definitions []Definition `json:"definitions"`
}

func (g Group) Get(ctx *gin.Context) {
	var request struct {
		Type string
	}

	var response struct {
		api.ResponseBase
		Word Word `json:"word"`
	}

	request.Type = ctx.Request.URL.Query().Get("type")

	wordID, wordSightID, err := getTodayID(ctx)
	if err != nil {
		code := api.DatabaseError
		logger.Error(code.Dump("Get today id error: %v", err))
		ctx.JSON(http.StatusInternalServerError, code.Response())
		return
	}

	collection := db.Collection(model.CWord).Where(english.Word_Key_ID.Eq(wordID))

	if request.Type == english.Type_Sight {
		var wordSight struct {
			ID              primitive.ObjectID `bson:"_id"`
			model.WordSight `bson:"-,inline"`
		}
		if err := db.Collection(model.CWordSight).
			Where(english.Word_Key_ID.Eq(wordSightID)).
			FindOne(
				ctx,
				&wordSight,
			); err != nil {
			code := api.DatabaseError
			logger.Error(code.Dump("Find word sight error: %v", err))
			ctx.JSON(http.StatusInternalServerError, code.Response())
			return
		}

		collection = collection.Where(english.Word_Key_Letters.Eq(wordSight.Letters))
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

func getTodayID(ctx *gin.Context) (primitive.ObjectID, primitive.ObjectID, error) {
	var wordToday struct {
		ID              primitive.ObjectID `bson:"_id"`
		model.WordToday `bson:"-,inline"`
	}

	if err := db.Collection(model.CWordToday).FindOneOrZero(ctx, &wordToday); err != nil {
		return primitive.NilObjectID, primitive.NilObjectID, err
	}

	if wordToday.ID.IsZero() {
		wordToday.ID = primitive.NewObjectID()
	}

	if wordToday.Date < time.Now().Format("2006-01-02") {
		var word struct {
			ID         primitive.ObjectID `bson:"_id"`
			model.Word `bson:"-,inline"`
		}

		if err := db.Collection(model.CWord).
			Where(english.Word_Key_ID.Gt(wordToday.WordID)).
			Sort(english.Word_Key_ID.Asc()).
			FindOneOrZero(
				ctx,
				&word,
			); err != nil {
			return primitive.NilObjectID, primitive.NilObjectID, err
		}
		if word.ID.IsZero() {
			if err := db.Collection(model.CWord).
				Sort(english.Word_Key_ID.Asc()).
				FindOneOrZero(
					ctx,
					&word,
				); err != nil {
				return primitive.NilObjectID, primitive.NilObjectID, err
			}
		}

		wordToday.WordID = word.ID

		var wordSight struct {
			ID              primitive.ObjectID `bson:"_id"`
			model.WordSight `bson:"-,inline"`
		}

		if err := db.Collection(model.CWordSight).
			Where(english.Word_Key_ID.Gt(wordToday.WordSightID)).
			Sort(english.Word_Key_ID.Asc()).
			FindOneOrZero(
				ctx,
				&wordSight,
			); err != nil {
			return primitive.NilObjectID, primitive.NilObjectID, err
		}

		if wordSight.ID.IsZero() {
			if err := db.Collection(model.CWordSight).
				Sort(english.Word_Key_ID.Asc()).
				FindOneOrZero(
					ctx,
					&wordSight,
				); err != nil {
				return primitive.NilObjectID, primitive.NilObjectID, err
			}
		}

		wordToday.WordSightID = wordSight.ID

		wordToday.Date = time.Now().Format("2006-01-02")

		if err := db.Collection(model.CWordToday).
			Where(english.Word_Key_ID.Eq(wordToday.ID)).
			Upsert(
				ctx,
				wordToday,
			); err != nil {
			return primitive.NilObjectID, primitive.NilObjectID, err
		}
	}

	return wordToday.WordID, wordToday.WordSightID, nil
}
