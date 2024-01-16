package models

type Ingredient struct {
	Article         int    `json:"article" db:"article"`
	Name            string `json:"name" db:"name"`
	MeasurementUnit string `json:"measurement_unit" db:"measurement_unit"`
	Count           int    `json:"count" db:"count"`
	MainSupplier    string `json:"main_supplier" db:"main_supplier"`
	Image           string `json:"image" db:"image"`
	Type            string `json:"type" db:"type"`
	PricePerUnit    string `json:"price_per_unit" db:"price_per_unit"`
	StateStandard   string `json:"state_standard" db:"state_standard"`
	PrePacking      string `json:"pre_packing" db:"pre_packing"`
	Characteristic  string `json:"characteristic" db:"characteristic"`
}
