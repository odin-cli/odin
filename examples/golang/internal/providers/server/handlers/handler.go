package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	logger *zap.Logger
}

func NewHandler(logger *zap.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) Create(c *gin.Context) {
	h.logger.Info("Called")
	c.Status(http.StatusCreated)
}

func (h *Handler) Get(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *Handler) Delete(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *Handler) Update(c *gin.Context) {
	c.Status(http.StatusOK)
}
