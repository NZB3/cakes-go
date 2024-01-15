package models

import "time"

type Suppliers struct {
	Name         string        `json:"name"`
	Address      string        `json:"address"`
	DeliveryTime time.Duration `json:"delivery_time"`
}
