package qrtickets

import (
	"time"
)

// Order - A PayPal ticket purchase order
type Order struct {
	Token          string    `json:"paypal_token" datastore:",noindex"`
	FirstName      string    `json:"firstname" datastore:",noindex"`
	LastName       string    `json:"lastname" datastore:",noindex"`
	Country        string    `json:"country" datastore:",noindex"`
	Amount         float32   `json:"amount" datastore:",noindex"`
	Currency       string    `json:"currency" datastore:",noindex"`
	ItemAmount     float32   `json:"item_amount" datastore:",noindex"`
	ShippingAmount float32   `json:"shipping_amount" datastore:",noindex"`
	TaxAmount      float32   `json:"tax_amount" datastore:",noindex"`
	Custom         string    `json:"custom" datastore:",noindex"`
	TxnID          string    `json:"txn_id"`
	Fees           float32   `json:"fees" datastore:",noindex"`
	OrderDate      time.Time `json:"order_date"`
	Email          string    `json:"email"`
}
