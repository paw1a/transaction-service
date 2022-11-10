package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/transaction-service/internal/domain"
	"github.com/paw1a/transaction-service/internal/domain/dto"
	"net/http"
	"strconv"
)

func (h *Handler) clientFindAll(context *gin.Context) {
	clients, err := h.clientService.FindAll(context.Request.Context())
	if err != nil {
		internalErrorResponse(context, err)
		return
	}

	clientsArray := make([]domain.Client, len(clients))
	if clients != nil {
		clientsArray = clients
	}

	successResponse(context, clientsArray)
}

func (h *Handler) clientFindById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		badRequestResponse(context,
			fmt.Sprintf("invalid client id param"), err)
		return
	}

	client, err := h.clientService.FindByID(context.Request.Context(), id)
	if err != nil {
		internalErrorResponse(context, fmt.Errorf("client with id: %d not found", id))
		return
	}

	successResponse(context, client)
}

func (h *Handler) clientCreate(context *gin.Context) {
	var clientDto dto.CreateClientDto
	err := context.BindJSON(&clientDto)
	if err != nil {
		badRequestResponse(context, "invalid client body format", err)
		return
	}

	client, err := h.clientService.Create(context.Request.Context(), clientDto)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	createdResponse(context, client)
}

func (h *Handler) clientDelete(context *gin.Context) {
	clientId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		errorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = h.clientService.Delete(context.Request.Context(), clientId)
	if err != nil {
		errorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.Status(http.StatusOK)
}
