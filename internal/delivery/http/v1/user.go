package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nekruzrabiev/simple-app/internal/domain"
	"github.com/nekruzrabiev/simple-app/internal/service"
	"net/http"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.POST("", h.userCreate)
		user.POST("/refresh", h.userRefresh)
		user.POST("/sign-in", h.userSignIn)
		authorized := user.Group("", h.userIdentity)
		{
			authorized.PUT("", h.userUpdateName)
			authorized.GET("", h.userGet)
		}
	}
}

func (h *Handler) userGet(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, ErrInternalServer, err)
		return
	}

	user, err := h.services.User.Get(c.Request.Context(), id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, ErrInternalServer, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

type userUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (h *Handler) userUpdateName(c *gin.Context) {
	var req userUpdateRequest
	if err := c.BindJSON(&req); err != nil {
		newResponse(c, http.StatusBadRequest, ErrBadParams, err)
		return
	}

	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, ErrInternalServer, err)
		return
	}

	err = h.services.User.UpdateName(c.Request.Context(), service.UserUpdateInput{
		Id:   id,
		Name: req.Name,
	})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, ErrInternalServer, err)
		return
	}

	c.Status(http.StatusOK)
}

type userSignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type signInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *Handler) userSignIn(c *gin.Context) {
	var req userSignInRequest
	if err := c.BindJSON(&req); err != nil {
		newResponse(c, http.StatusBadRequest, ErrBadParams, err)
		return
	}

	signInInfo, err := h.services.User.SignIn(c.Request.Context(), service.UserSignInInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, service.ErrEmailOrPassword) {
			newResponse(c, http.StatusBadRequest, ErrBadEmailOrPassword, err)
			return
		}
		newResponse(c, http.StatusInternalServerError, ErrInternalServer, err)
		return
	}

	c.JSON(http.StatusOK, signInResponse{
		AccessToken:  signInInfo.AccessToken,
		RefreshToken: signInInfo.RefreshToken,
	})
}

type userCreateRequest struct {
	FullName string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) userCreate(c *gin.Context) {
	var req userCreateRequest
	if err := c.BindJSON(&req); err != nil {
		newResponse(c, http.StatusBadRequest, ErrBadParams, err)
		return
	}

	id, err := h.services.User.Create(c.Request.Context(), domain.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, service.ErrInvalidPassword) {
			newResponse(c, http.StatusInternalServerError, ErrNotContainsDigitAndLetters, err)
			return
		}
		if errors.Is(err, service.ErrUserExists) {
			newResponse(c, http.StatusInternalServerError, ErrUserExists, err)
			return
		}
		newResponse(c, http.StatusInternalServerError, ErrInternalServer, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

type userRefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type userRefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *Handler) userRefresh(c *gin.Context) {
	var req userRefreshRequest
	if err := c.BindJSON(&req); err != nil {
		newResponse(c, http.StatusBadRequest, ErrBadParams, err)
		return
	}
	tokens, err := h.services.RefreshSession.Update(c.Request.Context(), req.RefreshToken)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, ErrSession, err)
		return
	}

	c.JSON(http.StatusOK, userRefreshResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}
