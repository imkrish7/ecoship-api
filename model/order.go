package model

import "time"

type Order struct {
	Id          string    `gorm:"type:uuid;primaryKey" json:"id"`
	OrderDate   time.Time `gorm:"type:timestamp" json:"order_date"`
	BuyerId     string    `gorm:"type:uuid" json:"buyer_id"`
	Buyer       Buyer     `gorm:"foreignKey:BuyerId"`
	Status      string    `gorm:"type:text" json:"status"`
	TotalAmount float64   `json:"total_amount"`
	ShippingId  string    `gorm:"type:uuid" json:"shipping_id"`
	Shipping    Shipping  `gorm:"foreignKey:ShippingId"`
	PaymentId   string    `gorm:"type:uuid" json:"payment_id"`
	Payment     Payment   `gorm:"foreignKey:PaymentId"`
	CreatedAt   time.Time `gorm:"type:timestamp" json:"created_at"`
	OrderItems  []OrderItem
}

type OrderItem struct {
	Id        string    `gorm:"type:uuid;primaryKey" json:"id"`
	OrderId   string    `json:"order_id"`
	ProductId string    `gorm:"type:uuid" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductId"`
	Quantity  int16     `json:"quantity"`
	InCart    bool      `gorm:"type:boolean" json:"in_cart"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
}

type Payment struct {
	Id        string    `grom:"type:uuid; primaryKey" json:"id"`
	OrderId   string    `json:"order_id"`
	Method    string    `gorm:"type:text" json:"method"`
	Amount    string    `gorm:"type:text" json:"amount"`
	Status    string    `gorm:"type:text" json:"status"`
	PaidAt    time.Time `gorm:"type:timestamp" json:"paid_at"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
}

type Shipping struct {
	Id             string    `gorm:"type:uuid;primaryKey" json:"id"`
	Address        string    `gorm:"type:text" json:"address"`
	City           string    `gorm:"type:text" json:"city"`
	PostalCode     string    `gorm:"type:text" json:"postal_code"`
	Country        string    `gorm:"type:text" json:"Country"`
	DeliveryStatus string    `gorm:"type:text" json:"delivery_status"`
	ShippedAt      time.Time `gorm:"type:timestamp" json:"shipped_at"`
	OrderId        string    `json:"order_id"`
	CreatedAt      time.Time `gorm:"type:timestamp" json:"created_at"`
}
