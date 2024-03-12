package main

import (
	"math"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func processReceipt(c *gin.Context) {
	var newReceipt Receipt

	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}
	uuid := uuid.New().String()
	receipts[uuid] = newReceipt
	c.IndentedJSON(http.StatusCreated, uuid)
}

// One point for every alphanumeric character in the retailer name
func getAlphanumericPoints(name string) int64 {
	var count int64 = 0
	for _, char := range name {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			count++
		}
	}
	return count
}

// 50 points if the total is a round dollar amount with no cents
func getRoundPoints(total float64) int64 {
	if total == float64(int64(total)) {
		return 50
	}
	return 0
}

// 25 points if the total is a multiple of 0.25
func getMultiplePoints(total float64) int64 {
	if math.Mod(total, 0.25) == 0 {
		return 25
	}
	return 0
}

// 5 points for every two items on the receipt
func getItemPoints(itemsLen int) int64 {
	return 5 * int64(itemsLen/2)
}

// 0.2 * price if description lenght is divisible by 3
func getItemDescPoints(items []Item) int64 {
	res := 0.0
	for i := 0; i < len(items); i++ {
		desc := strings.TrimSpace(items[i].ShortDescription)
		if len(desc)%3 == 0 {
			res += math.Ceil(0.2 * items[i].Price)
		}
	}
	return int64(res)
}

// 6 points if the day in the purchase date is odd
func getDatePoints(pDate time.Time) int64 {
	_, _, day := pDate.Date()
	if day%2 != 0 {
		return 6
	}
	return 0
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm
func getTimePoints(pTime time.Time) int64 {
	layout := "15:04"
	start, _ := time.Parse(layout, "14:00")
	end, _ := time.Parse(layout, "16:00")

	if pTime.After(start) && pTime.Before(end) {
		return 10
	}
	return 0
}

func getPoints(id string) int64 {
	var points int64 = 0
	receipt := receipts[id]
	points += getAlphanumericPoints(receipt.Retailer)
	points += getRoundPoints(receipt.Total)
	points += getMultiplePoints(receipt.Total)
	points += getItemPoints(len(receipt.Items))
	points += getItemDescPoints(receipt.Items)
	points += getDatePoints(receipt.PurchaseDate)
	points += getTimePoints(receipt.PurchaseTime)
	return points
}

func getPointsByID(c *gin.Context) {
	id := c.Param("id")
	points := getPoints(id)
	c.IndentedJSON(http.StatusOK, points)
}

var receipts = map[string]Receipt{}

func main() {
	r := gin.Default()
	r.POST("/receipts/process", processReceipt)
	r.GET("/receipts/:id/points", getPointsByID)

	r.Run()
}
