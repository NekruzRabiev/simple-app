package http

import (
	"github.com/gin-gonic/gin"
	_ "github.com/nekruzrabiev/simple-app/docs"
	v1 "github.com/nekruzrabiev/simple-app/internal/delivery/http/v1"
	"github.com/nekruzrabiev/simple-app/internal/service"
	"github.com/nekruzrabiev/simple-app/pkg/jwt"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"net/http"
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

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.tokenManager)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
