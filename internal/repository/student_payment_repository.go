package repository

import (
	"database/sql"
	"time"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/rs/zerolog/log"
)

type StudentPaymentRepository interface {
	CreatePayment(payment *model.StudentPayment) error
	GetPaymentByOrderID(orderID string) (*model.StudentPayment, error)
	UpdatePaymentStatus(orderID, status, transactionID string, paidAt *time.Time) error
}

type studentPaymentRepository struct {
	db *sql.DB
}

// NewStudentPaymentRepository creates a new instance of StudentPaymentRepository
func NewStudentPaymentRepository(db *sql.DB) StudentPaymentRepository {
	return &studentPaymentRepository{db: db}
}

// CreatePayment inserts a new payment record into the database
func (r *studentPaymentRepository) CreatePayment(payment *model.StudentPayment) error {
	query := `
		INSERT INTO student_payments (student_id, order_id, transaction_id, payment_method, payment_status, amount, snap_url, paid_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRow(
		query,
		payment.StudentID,
		payment.OrderID,
		payment.TransactionID,
		payment.PaymentMethod,
		payment.PaymentStatus,
		payment.Amount,
		payment.SnapURL,
		payment.PaidAt,
	).Scan(&payment.ID, &payment.CreatedAt, &payment.UpdatedAt)

	if err != nil {
		log.Error().Err(err).Msg("Failed to create payment record")
		return err
	}
	return nil
}

// GetPaymentByOrderID retrieves a payment record using the order ID
func (r *studentPaymentRepository) GetPaymentByOrderID(orderID string) (*model.StudentPayment, error) {
	query := `
		SELECT id, student_id, order_id, transaction_id, payment_method, payment_status, amount, snap_url, paid_at, created_at, updated_at
		FROM student_payments
		WHERE order_id = $1
	`
	payment := &model.StudentPayment{}
	err := r.db.QueryRow(query, orderID).Scan(
		&payment.ID,
		&payment.StudentID,
		&payment.OrderID,
		&payment.TransactionID,
		&payment.PaymentMethod,
		&payment.PaymentStatus,
		&payment.Amount,
		&payment.SnapURL,
		&payment.PaidAt,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Error().Err(err).Msg("Failed to retrieve payment record")
		return nil, err
	}
	return payment, nil
}

// UpdatePaymentStatus updates the payment status and optionally sets the paid_at timestamp
func (r *studentPaymentRepository) UpdatePaymentStatus(orderID, status, transactionID string, paidAt *time.Time) error {
	query := `
		UPDATE student_payments 
		SET payment_status = $1, transaction_id = $2, paid_at = $3, updated_at = CURRENT_TIMESTAMP
		WHERE order_id = $4
	`
	_, err := r.db.Exec(query, status, transactionID, paidAt, orderID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update payment status")
		return err
	}
	return nil
}

