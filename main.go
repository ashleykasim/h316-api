package main

import (
	"io"
	"log"
	"os"

	"github.com/Graylog2/go-gelf/gelf"
	"github.com/gin-gonic/gin"

	"recipes-api/config"
	"recipes-api/db"
	"recipes-api/models"
	"recipes-api/middleware"
)

var router *gin.Engine

func init() {
	config.ReadConfig("./config/config.toml")
}

func main() {
	graylogAddr := config.Get().GraylogAddr

	if graylogAddr != "" {
		gelfWriter, err := gelf.NewWriter(graylogAddr)
		if err != nil {
			log.Fatalf("gelf.NewWriter: %s", err)
		}
		// log to both stderr and graylog
		log.SetOutput(io.MultiWriter(os.Stderr, gelfWriter))
		log.Printf("logging to stderr & graylog2@'%s'", graylogAddr)
	} else {
		log.Printf("No GraylogAddr configured, logging to stderr only.")
	}

	db.Connect(config.Get().Ip, config.Get().Database)

  router = gin.Default()
	router.Use(middleware.CORSMiddleware)

	router.POST("/recipes", models.CreateRecipe)
	router.GET("/recipes", models.GetAllRecipes)
	router.GET("/recipes/:id", models.GetRecipeById)
	//TODO: get by name, not db id
	router.PUT("/recipes/:id", models.UpdateRecipe)
	router.DELETE("/recipes/:id", models.DeleteRecipe)

	router.GET("/menus", models.GetAllMenus)
	router.POST("/menus", models.CreateMenu)

	router.GET("/files/:name", models.DownloadFile)
	router.POST("/files", models.UploadFile)

	router.Run()
}
