package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// set the router as the default one provided by Gin
	router = gin.Default()

	// process the templates at the start of so that they dont have to be loaded
	// from the disk again. This makes serving HTML pages very fast.

	router.LoadHTMLGlob("templates/*")

	// define the route for the index page and display the index.html template
	// to start with, we'll use an inline route handler.
	router.GET("/", showIndexPage)
	// start serving the application
	router.Run()
}
