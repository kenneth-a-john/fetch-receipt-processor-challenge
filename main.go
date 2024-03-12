package main

import (
	"github.com/gin-gonic/gin"
)

var receipts = map[string]Receipt{}

func main() {
	r := gin.Default()
	r.POST("/receipts/process", ProcessReceipt)
	r.GET("/receipts/:id/points", GetPointsByID)

	r.Run()
}
