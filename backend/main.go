package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var r *gin.Engine
	r = gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/home", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"home.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})
	r.Run(":5000")
}
