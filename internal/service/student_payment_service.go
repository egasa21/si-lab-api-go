package service

import (
	"fmt"
	"time"

	"github.com/egasa21/si-lab-api-go/internal/model"
	"github.com/egasa21/si-lab-api-go/internal/repository"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"

	"github.com/rs/zerolog/log"
)

type StudentPaymentService interface {
	CreatePayment(studentID int, amount float64) (*model.StudentPayment, error)
	HandlePaymentNotification(orderID, transactionID, status string, paidAt *time.Time) error
	GetPaymentByOrderID(orderID string) (*model.StudentPayment, error)
}

type studentPaymentService struct {
	repo        repository.StudentPaymentRepository
	snapClient  snap.Client
	serverKey   string
	midtransEnv midtrans.EnvironmentType
}

// NewStudentPaymentService initializes the payment service with Midtrans Snap
func NewStudentPaymentService(repo repository.StudentPaymentRepository, serverKey string, midtransEnv midtrans.EnvironmentType) StudentPaymentService {
	snapClient := snap.Client{}
	snapClient.New(serverKey, midtransEnv)

	return &studentPaymentService{
		repo:        repo,
		snapClient:  snapClient,
		serverKey:   serverKey,
		midtransEnv: midtransEnv,
	}
}

// CreatePayment initiates a new Midtrans Snap transaction and stores payment details in DB
func (s *studentPaymentService) CreatePayment(studentID int, amount float64) (*model.StudentPayment, error) {
	orderID := fmt.Sprintf("ORDER-%d-%d", studentID, time.Now().Unix())

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	// Create transaction with Midtrans
	snapResp, err := s.snapClient.CreateTransaction(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create Midtrans transaction")
		return nil, err
	}

	// Store payment in DB
	payment := &model.StudentPayment{
		StudentID:     studentID,
		OrderID:       orderID,
		TransactionID: "", 
		PaymentMethod: "midtrans",
		PaymentStatus: "pending",
		Amount:        amount,
		SnapURL:       snapResp.RedirectURL,
		PaidAt:        nil, 
	}

	createErr := s.repo.CreatePayment(payment)
	if createErr != nil {
		log.Error().Err(createErr).Msg("Failed to store payment record")
		return nil, createErr
	}

	return payment, nil
}

// HandlePaymentNotification updates payment status upon Midtrans notification
func (s *studentPaymentService) HandlePaymentNotification(orderID, transactionID, status string, paidAt *time.Time) error {
	return s.repo.UpdatePaymentStatus(orderID, status, transactionID, paidAt)
}

// GetPaymentByOrderID retrieves a payment record by order ID
func (s *studentPaymentService) GetPaymentByOrderID(orderID string) (*model.StudentPayment, error) {
	return s.repo.GetPaymentByOrderID(orderID)
}

func (s *studentPaymentService) UpdatePaymentStatus(orderID, status, transactionID string, paidAt *time.Time) error {
	return s.repo.UpdatePaymentStatus(orderID, status, transactionID, paidAt)
}