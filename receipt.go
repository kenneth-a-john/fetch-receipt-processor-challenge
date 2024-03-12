package main

import (
	"encoding/json"
	"strconv"
	"time"
)

type Item struct {
	ShortDescription string  `json:"shortdescription"`
	Price            float64 `json:"price"`
}

type Receipt struct {
	Retailer     string    `json:"retailer"`
	PurchaseDate time.Time `json:"purchasedate"`
	PurchaseTime time.Time `json:"purchasetime"`
	Total        float64   `json:"total"`
	Items        []Item    `json:"items"`
}

type ItemAux struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type ReceiptAux struct {
	Retailer     string    `json:"retailer"`
	PurchaseDate string    `json:"purchaseDate"`
	PurchaseTime string    `json:"purchaseTime"`
	Total        string    `json:"total"`
	Items        []ItemAux `json:"items"`
}

func (r *Receipt) UnmarshalJSON(data []byte) error {
	var aux ReceiptAux
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Convert PurchaseDate and PurchaseTime
	purchaseDate, err := time.Parse("2006-01-02", aux.PurchaseDate)
	if err != nil {
		return err
	}

	purchaseTime, err := time.Parse("15:04", aux.PurchaseTime)
	if err != nil {
		return err
	}

	// Convert Total
	total, err := strconv.ParseFloat(aux.Total, 64)
	if err != nil {
		return err
	}

	r.Retailer = aux.Retailer
	r.PurchaseDate = purchaseDate
	r.PurchaseTime = purchaseTime
	r.Total = total

	// Convert each item in Items
	for _, itemAux := range aux.Items {
		price, err := strconv.ParseFloat(itemAux.Price, 64)
		if err != nil {
			return err
		}
		item := Item{
			ShortDescription: itemAux.ShortDescription,
			Price:            price,
		}
		r.Items = append(r.Items, item)
	}

	return nil
}
