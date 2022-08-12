package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type FormBody struct {
	Nama      string `json:"nama"`
	Reason    string `json:"reason"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

var excelFile = "overtime_data.xlsx"
var sheet = "Sheet1"

func formPost(c *gin.Context) {
	reqBody := FormBody{}
	if err := c.BindJSON(&reqBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	startDate := reqBody.StartDate

	fmt.Println(string(startDate))
	// for _, val := range startDate {
	// 	fmt.Println(string(val))
	// }
	// year := startDate[0:4]
	// month := startDate[6:8]
	// day := startDate[8:10]
	c.JSON(http.StatusAccepted, &reqBody)

	// Insert New Data
	fOpen, err := excelize.OpenFile(excelFile)
	defer func() {
		if err := fOpen.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		// if file not exist create and write excel
		fCreate := excelize.NewFile()
		defer func() {
			if err := fCreate.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		fCreate.SetCellValue(sheet, "A1", "No")
		fCreate.SetCellValue(sheet, "B1", "Nama")
		fCreate.SetCellValue(sheet, "C1", "Alasan")
		fCreate.SetCellValue(sheet, "D1", "Start")
		fCreate.SetCellValue(sheet, "E1", "End")

		if err := fCreate.SaveAs(excelFile); err != nil {
			log.Fatal(err)
		}
		fCreate.Close()
	}

	// get all rows
	rows, err := fOpen.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
		return
	}
	// get last row
	var lastRow int
	for i := 0; i <= len(rows); i++ {
		lastRow++
	}
	fOpen.SetActiveSheet(1)
	rowA := fmt.Sprintf("A%d", lastRow)
	rowB := fmt.Sprintf("B%d", lastRow)
	rowC := fmt.Sprintf("C%d", lastRow)
	rowD := fmt.Sprintf("D%d", lastRow)
	rowE := fmt.Sprintf("E%d", lastRow)
	fOpen.SetCellValue(sheet, rowA, lastRow)
	fOpen.SetCellValue(sheet, rowB, reqBody.Nama)
	fOpen.SetCellValue(sheet, rowC, reqBody.Reason)

	fOpen.SetCellValue(sheet, rowD, reqBody.StartDate)
	fOpen.SetCellValue(sheet, rowE, reqBody.EndDate)

	if err := fOpen.Save(); err != nil {
		log.Fatal(err)
	}
	// if err := fOpen.Close(); err != nil {
	// 	fmt.Println(err)
	// }

}

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

func main() {
	var r *gin.Engine
	r = gin.Default()

	r.Use(CORSMiddleware())
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"data": "It Works!",
			},
		)
	})
	r.POST("/form", formPost)

	// })
	r.Run(":5000")
}
