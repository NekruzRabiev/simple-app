package v1

import (
	"bytes"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/nekruzrabiev/simple-app/internal/domain"
	"github.com/nekruzrabiev/simple-app/internal/service"
	serviceMocks "github.com/nekruzrabiev/simple-app/internal/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_userCreate(t *testing.T) {
	type mockBehavior func (r *serviceMocks.MockUser, user domain.User)

	tests := []struct {
		name string
		inputBody string
		inputUser domain.User
		mockBehavior mockBehavior
		expectedStatusCode int
		expectedResponseBody string
	} {
		{
			name: "OK",
			inputBody: `{ "name": "Adam", "email": "adam@example.com", "password": "password123"}`,
			inputUser: domain.User{
				FullName: "Adam",
				Email: "adam@example.com",
				Password: "password123",
			},
			mockBehavior: func(r *serviceMocks.MockUser, user domain.User) {
				r.EXPECT().Create(context.Background(), user).Return(1, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name: "Wrong Input",
			inputBody: `{"name":"Adam"}`,
			inputUser: domain.User{},
			mockBehavior: func(r *serviceMocks.MockUser, user domain.User) {},
			expectedStatusCode: 400,
			expectedResponseBody: ErrBadParams.marshalJSON(),
		},
		{
			name: "Wrong Email",
			inputBody: `{ "name": "Adam", "email": "adam@example", "password": "password123"}`,
			inputUser: domain.User{},
			mockBehavior: func(r *serviceMocks.MockUser, user domain.User) {},
			expectedStatusCode: 400,
			expectedResponseBody: ErrBadParams.marshalJSON(),
		},
		{
			name: "Invalid Password",
			inputBody: `{ "name": "Adam", "email": "adam@example.com", "password": "password"}`,
			inputUser: domain.User{
				FullName: "Adam",
				Email: "adam@example.com",
				Password: "password",
			},
			mockBehavior: func(r *serviceMocks.MockUser, user domain.User) {
				r.EXPECT().Create(context.Background(), user).Return(0, service.ErrInvalidPassword)
			},
			expectedStatusCode: 400,
			expectedResponseBody: ErrNotContainsDigitAndLetters.marshalJSON(),
		},
		{
			name: "User Exists",
			inputBody: `{ "name": "Adam", "email": "adam@example.com", "password": "password"}`,
			inputUser: domain.User{
				FullName: "Adam",
				Email: "adam@example.com",
				Password: "password",
			},
			mockBehavior: func(r *serviceMocks.MockUser, user domain.User) {
				r.EXPECT().Create(context.Background(), user).Return(0, service.ErrUserExists)
			},
			expectedStatusCode: 500,
			expectedResponseBody: ErrUserExists.marshalJSON(),
		},
		{
			name: "Service Error",
			inputBody: `{ "name": "Adam", "email": "adam@example.com", "password": "password"}`,
			inputUser: domain.User{
				FullName: "Adam",
				Email: "adam@example.com",
				Password: "password",
			},
			mockBehavior: func(r *serviceMocks.MockUser, user domain.User) {
				r.EXPECT().Create(context.Background(), user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode: 500,
			expectedResponseBody: ErrInternalServer.marshalJSON(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := serviceMocks.NewMockUser(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Services{User: repo}
			handler := Handler{services: services}

			// init endpoint
			r := gin.New()
			r.POST("/", handler.userCreate)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}


