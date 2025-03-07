package handler

import (
	"encoding/json"
	"net/http"

	"github.com/egasa21/si-lab-api-go/internal/dto"
	"github.com/egasa21/si-lab-api-go/internal/pkg"
	"github.com/egasa21/si-lab-api-go/internal/pkg/response"
	"github.com/egasa21/si-lab-api-go/internal/service"
)

type StudentPaymentHandler struct {
	service service.StudentPaymentService
}

func NewStudentPaymentHandler(service service.StudentPaymentService) *StudentPaymentHandler {
	return &StudentPaymentHandler{service: service}
}

// CreatePayment handles creating a new payment
func (h *StudentPaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	payment, err := h.service.CreatePayment(req.StudentID, req.Amount)
	if err != nil {
		appErr := pkg.NewAppError("Failed to create payment", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	resp := dto.CreatePaymentResponse{
		OrderID:       payment.OrderID,
		TransactionID: payment.TransactionID,
		Amount:        payment.Amount,
		SnapURL:       payment.SnapURL,
		PaymentStatus: payment.PaymentStatus,
	}

	response.NewSuccessResponse(w, resp, "Payment created successfully")
}

// HandlePaymentNotification processes incoming notifications from Midtrans
func (h *StudentPaymentHandler) HandlePaymentNotification(w http.ResponseWriter, r *http.Request) {
	var req dto.PaymentNotificationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		appErr := pkg.NewAppError("Invalid request payload", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	// Handle the notification by updating payment status
	err := h.service.HandlePaymentNotification(req.OrderID, req.TransactionID, req.Status, req.PaidAt)
	if err != nil {
		appErr := pkg.NewAppError("Failed to handle payment notification", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, nil, "Payment status updated successfully")
}

// GetPayment retrieves payment details by order ID
func (h *StudentPaymentHandler) GetPaymentByOrderID(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("order_id")
	if orderID == "" {
		appErr := pkg.NewAppError("Order ID is required", http.StatusBadRequest)
		response.NewErrorResponse(w, appErr)
		return
	}

	payment, err := h.service.GetPaymentByOrderID(orderID)
	if err != nil {
		appErr := pkg.NewAppError("Failed to retrieve payment", http.StatusInternalServerError)
		response.NewErrorResponse(w, appErr)
		return
	}

	if payment == nil {
		appErr := pkg.NewAppError("Payment not found", http.StatusNotFound)
		response.NewErrorResponse(w, appErr)
		return
	}

	response.NewSuccessResponse(w, payment, "Payment retrieved successfully")
}
