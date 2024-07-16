package models

type Amount struct {
	Amount float64 `json:"Amount" binding:"required"`
}
