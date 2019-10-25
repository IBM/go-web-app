package main

import (
	routers "github.com/IBM/go-web-app/routers"
	// "gowebapp/plugins" if you create your own plugins import them here
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	router := gin.Default()
	router.RedirectTrailingSlash = false

	router.LoadHTMLGlob("public/*.html")
	router.Use(static.Serve("/", static.LocalFile("./public", false)))
	router.GET("/", routers.Index)
	router.NoRoute(routers.NotFoundError)
	router.GET("/500", routers.InternalServerError)
	router.GET("/health", routers.HealthGET)

	log.Info("Starting gowebapp on port " + port())

	router.Run(port())
}
