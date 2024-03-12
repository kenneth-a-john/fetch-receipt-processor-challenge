package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProcessReceipt(c *gin.Context) {
	var newReceipt Receipt

	if err := c.BindJSON(&newReceipt); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Input does not match with definition!"})
		return
	}
	uuid := uuid.New().String()
	receipts[uuid] = newReceipt
	c.IndentedJSON(http.StatusCreated, uuid)
}
