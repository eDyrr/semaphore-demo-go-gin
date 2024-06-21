package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	render(c, gin.H{"title": "Home Page", "payload": articles}, "index.html")
}

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// respond with json
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// respond with xml
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

func getArticle(c *gin.Context) {
	// check if the passed ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// checl if article exists
		if article, err := getArticleByID(articleID); err == nil {
			render(c, gin.H{"title": article.Title, "payload": article}, "article.html")
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
