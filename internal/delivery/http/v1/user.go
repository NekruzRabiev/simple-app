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
			authorized.PATCH("", h.userUpdateName)
			authorized.GET("", h.userGet)
		}
	}
}

// @Summary userGet
// @Security UsersAuth
// @Tags users
// @Description Get user's data
// @ID get-user
// @Produce  json
// @Success 200 {object} domain.User
// @Failure 400,404,401 {object} errResponse
// @Failure 500 {object} errResponse
// @Failure default {object} errResponse
// @Router /api/v1/user [get]
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

// @Summary userUpdateName
// @Security UsersAuth
// @Tags users
// @Description Update user's name
// @ID user-update-name
// @Accept  json
// @Param input body userUpdateRequest true "Name for updating"
// @Success 200
// @Failure 400,404,401 {object} errResponse
// @Failure 500 {object} errResponse
// @Failure default {object} errResponse
// @Router /api/v1/user [patch]
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

// @Summary SignIn
// @Tags users
// @Description Sign in
// @ID sign-in
// @Accept  json
// @Produce  json
// @Param input body userSignInRequest true "account info"
// @Success 200 {object} signInResponse
// @Failure 400,404 {object} errResponse
// @Failure 500 {object} errResponse
// @Failure default {object} errResponse
// @Router /api/v1/user/sign-in [post]
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

// @Summary userCreate
// @Tags users
// @Description Create new user
// @ID create-user
// @Accept  json
// @Produce  json
// @Param input body userCreateRequest true "user's info"
// @Success 200 integer id "returns id"
// @Failure 400,404 {object} errResponse
// @Failure 500 {object} errResponse
// @Failure default {object} errResponse
// @Router /api/v1/user [post]
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

// @Summary userRefresh
// @Tags users
// @Description Refresh user's tokens
// @ID refresh-user
// @Accept  json
// @Produce  json
// @Param input body userRefreshRequest true "refresh token"
// @Success 200 {object} userRefreshResponse
// @Failure 400,404 {object} errResponse
// @Failure 500 {object} errResponse
// @Failure default {object} errResponse
// @Router /api/v1/user/refresh [post]
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
