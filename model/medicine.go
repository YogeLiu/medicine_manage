package model

type Medicine struct {
	ID     string  `json:"id"` // id equal name
	Name   string  `json:"name"`
	Amount int     `json:"amount"`
	Price  float32 `json:"price"`
}

type MedicineList struct {
	Medicines []*Medicine `json:"medicines"`
	HasMore   bool        `json:"has_more"`
}
