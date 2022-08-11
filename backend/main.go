package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
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

var sheet = "overtime_data"

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
	fCreate := excelize.NewFile()
	// new sheet
	index := fCreate.NewSheet(sheet)
	fCreate.SetCellValue(sheet, "A1", "No")
	fCreate.SetCellValue(sheet, "B1", "Nama")
	fCreate.SetCellValue(sheet, "C1", "Alasan")
	fCreate.SetCellValue(sheet, "D1", "Start")
	fCreate.SetCellValue(sheet, "E1", "End")
	fCreate.SetActiveSheet(index)
	// now := time.Now()
	// f.SetCellValue("Sheet1", "A4", now.Format(time.ANSIC))
	if err := fCreate.SaveAs(sheet + ".xlsx"); err != nil {
		log.Fatal(err)
	}

	// // get rows
	// fOpen, err := excelize.OpenFile(sheet)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer func() {
	// 	if err := fOpen.Close(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()
	// rows, err := fOpen.GetCellValue(sheet, "A1")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(rows)

}

func main() {
	var r *gin.Engine
	r = gin.Default()

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
	r.GET("/getExc", func(ctx *gin.Context) {
		// get rows
		fOpen, err := excelize.OpenFile(sheet + ".xlsx")
		if err != nil {
			fmt.Println(err)
			return
		}
		// defer func() {
		// 	if err := fOpen.Close(); err != nil {
		// 		fmt.Println(err)
		// 	}
		// }()
		rows, err := fOpen.GetCellValue(sheet, "A1")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("================")
		fmt.Println(rows)
		fmt.Println("================")

	})
	r.Run(":5000")
}
