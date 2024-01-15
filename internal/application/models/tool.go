package models

import "time"

type Tool struct {
	Name           string    `json:"name" db:"name"`
	Description    string    `json:"description" db:"description"`
	Equipment      string    `json:"equipment" db:"equipment"`
	WearRate       int       `json:"wear_rate" db:"wear_rate"`
	Supplier       string    `json:"supplier" db:"supplier"`
	DateOfPurchase time.Time `json:"date_of_purchase" db:"date_of_purchase"`
	Count          int       `json:"count" db:"count"`
}
