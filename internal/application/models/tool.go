package models

import "time"

type Tool struct {
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Equipment      string    `json:"equipment"`
	WearRate       int       `json:"wear_rate"`
	Supplier       string    `json:"supplier"`
	DateOfPurchase time.Time `json:"date_of_purchase"`
	Count          int       `json:"count"`
}
