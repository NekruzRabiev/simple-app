package v1

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.POST("/sign-in", h.create)
		user.POST("/refresh", h.userRefresh)
		authorized := user.Group("", h.userIdentity)
		{
			authorized.PUT("", h.userUpdate)
			authorized.GET("", h.userGet)
		}
	}
}