package api

import (
	"encoding-restoration/internal/controllers"
	"encoding-restoration/internal/models"
	"log"

	"github.com/gin-gonic/gin"
)

var DEBUG bool = false

func StartServer() {

	log.Println("Server start up")

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/errors", controllers.ErrorPage)

	r.GET("/syndromes", controllers.SyndromePage)

	r.GET("/syndromes/array", controllers.SyndromeArrayPage)

	r.GET("/results/", controllers.ResultsPage)

	r.Static("/assets", "./resources")

	models.Calculate()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
