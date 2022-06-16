package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thnkrn/go-gin-template/pkg/usecase"
)

type MockHandler struct {
	usecase usecase.Usecase
}

func NewHandler(usecase usecase.Usecase) *MockHandler {
	return &MockHandler{
		usecase: usecase,
	}
}

func (h *MockHandler) Mock(c *gin.Context) {
	txt := h.usecase.Mock(c.Request.Context())

	c.JSON(http.StatusOK, txt)
}
