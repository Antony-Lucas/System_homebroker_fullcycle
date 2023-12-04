package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Price        float64
	Total        float64
	DateTime     time.Time
}

func NewTransaction(SellingOrder *Order, BuyingOrder *Order, shares int, price float64) *Transaction{
	total := float64(shares) * price
	return &Transaction{
		ID: uuid.New().String(),
		SellingOrder: SellingOrder,
		BuyingOrder: BuyingOrder,
		Shares: shares,
		Price: price,
		Total: total,
		DateTime: time.Now(),
	}
}

func (t *Transaction) CalculateTotal(shares int, price float64){
	t.Total = float64(t.Shares) * t.Price
}

func (t *Transaction) CloseBuyOrder(){
	if t.BuyingOrder.PendingShares == 0 {
		t.BuyingOrder.Stats = "CLOSED"
	} 
}

func (t *Transaction) CloseSellOrder(){
	if t.SellingOrder.PendingShares == 0 {
		t.SellingOrder.Stats = "CLOSED"
	}
}

func (t *Transaction) AddBuyPendingShares(shares int){
	t.SellingOrder.PendingShares += shares
}

func (t *Transaction) AddSellPendingShares(shares int){
	t.BuyingOrder.PendingShares += shares
}