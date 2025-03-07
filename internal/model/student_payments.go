package model

import (
	"time"
)

type StudentPayment struct {
	ID            int       `json:"id"`
	StudentID     int       `json:"student_id"`
	OrderID       string    `json:"order_id"`
	TransactionID string    `json:"transaction_id"`
	PaymentMethod string    `json:"payment_method"`
	PaymentStatus string    `json:"payment_status" gorm:"default:pending"`
	Amount        float64   `json:"amount"`
	SnapURL       string    `json:"snap_url"`
	PaidAt        *time.Time `json:"paid_at,omitempty"` 
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
