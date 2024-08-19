package handler

import (
	"sj/internal/dto"
	"sj/internal/service"
	"sj/pkg/response"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(s service.TransactionService) *TransactionHandler {
	return &TransactionHandler{s}
}

func (h *TransactionHandler) GetAllByUserID(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(dto.UserClaims)

	transactions, err := h.service.GetAllByUserID(claims.ID)
	if err != nil {
		response.FailOrError(c, 400, "Failed getting transactions data", err)
		return
	}
	response.Success(c, 200, "Success getting user", transactions)
}
