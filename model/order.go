package model

type Order struct {
	ID       string  `json:"id"`
	Detail   []byte  `json:"detail"`
	Doctor   string  `json:"doctor"`
	Patient  string  `json:"patient"`
	Price    float32 `json:"price"`
	CreateAt int64   `json:"create_at"`
}
