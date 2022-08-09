package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func formPost(c *gin.Context) {
	c.Request.ParseForm()
	for key, val := range c.Request.PostForm {
		fmt.Println(key, val)
	}
}

func main() {
	var r *gin.Engine
	r = gin.Default()
	// r.LoadHTMLGlob("views/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"data": "It Works!",
			},
		)
	})
	r.POST("/", formPost)
	r.Run(":5000")
}
