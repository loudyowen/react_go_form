package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
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

type FormBody struct {
	Nama       string `json:"nama"`
	Reason     string `json:"reason"`
	StartValue string `json:"startValue"`
	EndValue   string `json:"endValue"`
}

func formPost(c *gin.Context) {
	fmt.Println("+===========START=============+")
	reqBody := FormBody{}
	if err := c.BindJSON(&reqBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println("Nama: ", reqBody.Nama)
	fmt.Println("Reason: ", reqBody.Reason)
	fmt.Println("Start : ", reqBody.StartValue)
	fmt.Println("End : ", reqBody.EndValue)
	c.JSON(http.StatusAccepted, &reqBody)

	fmt.Println("+============END==============+")

	// create and write excel
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "No")
	f.SetCellValue("Sheet1", "B1", "Nama")
	f.SetCellValue("Sheet1", "C1", "Alasan")
	f.SetCellValue("Sheet1", "D1", "Start")
	f.SetCellValue("Sheet1", "E1", "End")
	// now := time.Now()
	// f.SetCellValue("Sheet1", "A4", now.Format(time.ANSIC))
	if err := f.SaveAs("Simple.xlsx"); err != nil {
		log.Fatal(err)
	}
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
