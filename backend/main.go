package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	// "github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func formPost(c *gin.Context) {
	fmt.Println("+===========START=============+")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	fmt.Println("+============END==============+")
	// fmt.Println(c.Request.ParseForm())
	// for key, val := range c.Request.PostForm {
	// fmt.Println(key, val)
	// }
}

func main() {
	var r *gin.Engine
	r = gin.Default()
	// load everything inside a folder
	// r.LoadHTMLGlob("views/*")

	// r.Use(cors.New(cors.Config{
	// 	// AllowMethod: []string{"PUT", "PATCH"}
	// 	AllowOrigins:     []string{"http://127.0.0.1:5000"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	// AllowOriginFunc: func(origin string) bool{
	// 	// 	return origin == "http://github.com"
	// 	// },
	// 	MaxAge: 12 * time.Hour,
	// }))
	r.Use(CORSMiddleware())
	r.GET("/form", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"data": "It Works!",
			},
		)
	})
	r.POST("/form", formPost)
	r.Run(":5000")
}
