package entity

import "time"

type Transacion struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Price        float64
	Total        float64
	DateTime     time.Time
}

func NewTransaction(SellingOrder *Order, BuyingOrder *Order, shares int, price float64) *Transacion {
	total := float64(shares) * price
}
