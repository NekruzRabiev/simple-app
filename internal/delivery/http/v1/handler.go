package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nekruzrabiev/simple-app/internal/service"
	"github.com/nekruzrabiev/simple-app/pkg/jwt"
)

type Handler struct {
	services     *service.Services
	tokenManager jwt.TokenManager
}

func NewHandler(services *service.Services, tokenManager jwt.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initUserRoutes(v1)
	}
}
