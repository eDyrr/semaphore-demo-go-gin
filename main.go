package main

import (
	"net/http"

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
	router.GET("/", func(c *gin.Context) {

		// call the html method of the context to render a template
		c.HTML(
			// set the HTTP status to 200 (OK)
			http.StatusOK,
			// use the inedx.html template
			"index.html",
			// pass the data that the page uses (int this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)
	})
	// start serving the application
	defer router.Run()
}
