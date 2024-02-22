package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"hello-english/api"
	"hello-english/base/openai"
	"hello-english/db"

	"github.com/Lee-Chi/go-sdk/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Config struct {
	DBDomain   string `json:"mongodb_domain"`
	DBUser     string `json:"mongodb_user"`
	DBPassword string `json:"mongodb_password"`

	OpenaiToken string `json:"openai_token"`
}

func CorsConfig() cors.Config {
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
	corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	return corsConf
}

func main() {
	dbDomain := os.Getenv("MONGODB_DOMAIN")
	dbUser := os.Getenv("MONGODB_USER")
	dbPassword := os.Getenv("MONGODB_PASSWORD")
	openaiToken := os.Getenv("OPENAI_TOKEN")

	conf := flag.String("config", "", "")
	flag.Parse()

	if *conf != "" {
		data, err := os.ReadFile(*conf)
		if err != nil {
			panic(err)
		}

		config := Config{}
		if err := json.Unmarshal(data, &config); err != nil {
			panic(err)
		}

		dbDomain = config.DBDomain
		dbUser = config.DBUser
		dbPassword = config.DBPassword

		openaiToken = config.OpenaiToken
	}

	if err := openai.Build(openaiToken); err != nil {
		panic(err)
	}

	uri := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		dbUser,
		dbPassword,
		dbDomain,
	)

	if err := db.Build(context.Background(), uri, "helloenglishdb"); err != nil {
		panic(err)
	}

	logger.Init()

	logger.Info("service start ....")

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./out", false)))
	router.GET("/", func(ctx *gin.Context) {
		ctx.File("./out/index.html")
	})

	router.Use(cors.New(CorsConfig()))

	router.GET("/api/word", api.Word.Get)
	router.POST("/api/word/forward", api.Word.Forward)
	router.POST("/api/word/backward", api.Word.Backward)
	router.POST("/api/word/explain", api.Word.Explain)

	router.POST("/api/sentence/check", api.Sentence.Check)
	router.POST("/api/sentence/translate", api.Sentence.Translate)
	router.POST("/api/sentence/advise", api.Sentence.Advise)
	router.POST("/api/sentence/practice/ready", api.Sentence.Practice.Ready)
	router.POST("/api/sentence/practice/submit", api.Sentence.Practice.Submit)

	router.POST("/api/paragraph/explain", nil)
	router.POST("/api/paragraph/extract", nil)

	router.POST("/api/article/chat", nil)

	go func() {
		if err := router.Run(":8080"); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-quit

	if err := db.Destroy(context.Background()); err != nil {
		panic(err)
	}

	logger.Info("service stop ....")
}
