package models

import "time"

type Suppliers struct {
	Name         string        `json:"name" db:"name"`
	Address      string        `json:"address" db:"address"`
	DeliveryTime time.Duration `json:"delivery_time" db:"delivery_time"`
}
