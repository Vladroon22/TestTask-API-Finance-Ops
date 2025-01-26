package handlers

import (
	"net/http"
	"strconv"

	golog "github.com/Vladroon22/GoLog"
	"github.com/Vladroon22/TestTask-Bank-Operation/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Servicer
	logger  *golog.Logger
}

func NewHandler(service service.Servicer, lg *golog.Logger) *Handler {
	return &Handler{service: service, logger: lg}
}

func (h *Handler) IncreaseUserBalance(c *gin.Context) {
	var request struct {
		UserID int     `json:"user_id"`
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.logger.Errorln(err)
		return
	}

	err := h.service.IncreaseUserBalance(c.Request.Context(), request.UserID, request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logger.Errorln(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h *Handler) TransferMoney(c *gin.Context) {
	var request struct {
		fromUserID int     `json:"from_user_id"`
		toUserID   int     `json:"to_user_id"`
		userFrom   string  `json:"sender_name"`
		userTo     string  `json:"receiver_name"`
		amount     float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.logger.Errorln(err)
		return
	}

	if err := h.service.TransferMoney(c.Request.Context(), request.userFrom, request.userTo, request.fromUserID, request.toUserID, request.amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logger.Errorln(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Money has tranfered to ": request.userTo})
}

func (h *Handler) GetLastTxs(c *gin.Context) {
	id := c.Param("userID")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		h.logger.Errorln("Error: invalid user ID")
		return
	}

	transactions, err := h.service.GetLastTxs(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logger.Errorln(err)
		return
	}

	c.JSON(http.StatusOK, transactions)
}
