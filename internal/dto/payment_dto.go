package dto

import "time"

type CreatePaymentRequest struct {
	StudentID int     `json:"student_id"`
	Amount    float64 `json:"amount"`
}

type CreatePaymentResponse struct {
	OrderID      string  `json:"order_id"`
	TransactionID string `json:"transaction_id"`
	Amount       float64 `json:"amount"`
	SnapURL      string  `json:"snap_url"`
	PaymentStatus string `json:"payment_status"`
}

type PaymentNotificationRequest struct {
	OrderID      string     `json:"order_id"`
	TransactionID string     `json:"transaction_id"`
	Status       string     `json:"status"`
	PaidAt       *time.Time `json:"paid_at"`
}
