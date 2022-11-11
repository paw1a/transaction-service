package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/transaction-service/internal/domain"
	"github.com/paw1a/transaction-service/internal/domain/dto"
	"net/http"
	"strconv"
)

func (h *Handler) transactionFindAll(context *gin.Context) {
	transactions, err := h.transactionService.FindAll(context.Request.Context())
	if err != nil {
		internalErrorResponse(context, err)
		return
	}

	transactionsArray := make([]domain.Transaction, len(transactions))
	if transactions != nil {
		transactionsArray = transactions
	}

	successResponse(context, transactionsArray)
}

func (h *Handler) transactionFindById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		badRequestResponse(context,
			fmt.Sprintf("invalid transaction id param"), err)
		return
	}

	transaction, err := h.transactionService.FindById(context.Request.Context(), int64(id))
	if err != nil {
		internalErrorResponse(context, fmt.Errorf("transaction with id: %d not found", id))
		return
	}

	successResponse(context, transaction)
}

func (h *Handler) transactionCreate(context *gin.Context) {
	var transactionDto dto.CreateTransactionDto
	err := context.BindJSON(&transactionDto)
	if err != nil {
		badRequestResponse(context, "invalid transaction body format", err)
		return
	}

	transaction, err := h.transactionService.Create(context.Request.Context(), transactionDto)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	createdResponse(context, transaction)
}
