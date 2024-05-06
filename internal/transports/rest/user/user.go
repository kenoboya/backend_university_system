package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"test-crud/internal/model"
	"test-crud/internal/service"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type UsersHandler struct {
	usersService      service.Users
	complaintsService service.Complaints
}

func NewUsersHandler(usersService service.Users, complaintsService service.Complaints) *UsersHandler {
	return &UsersHandler{
		usersService:      usersService,
		complaintsService: complaintsService,
	}
}

// @Summary User registration
// @Description User registration
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.UserSignUpInput true "Data for registration user"
// @Success 201 {string} string "Registered"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /sign-up [post]
func (h *UsersHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "signUp()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user model.UserSignUpInput
	if err := json.Unmarshal(reqBytes, &user); err != nil {
		zap.S().Error(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "signUp()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		zap.S().Error(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "signUp()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.usersService.SignUp(context.TODO(), user); err != nil {
		zap.S().Error(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "signUp()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary Sign In User
// @Description Sign in a user and generate access token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.UserSignInInput true "Data for signing in user"
// @Success 200 {object} service.Tokens "Successful operation"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /sign-in [post]
func (h *UsersHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "signIn()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user model.UserSignInInput
	if err := json.Unmarshal(reqBytes, &user); err != nil {
		zap.S().Error(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "signIn()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		zap.S().Error(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "signIn()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token, err := h.usersService.SignIn(context.TODO(), user)
	if err != nil {
		zap.S().Error(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "signIn()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := json.Marshal(map[string]string{
		"token": token.AccessToken,
	})
	if err != nil {
		zap.S().Error(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "signIn()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", token.RefreshToken))
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

// @Summary Refresh Access Token
// @Description Refresh access token using refresh token stored in cookie
// @Tags auth
// @Produce json
// @Success 200 {object} service.Tokens "Successful operation"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /refresh [get]
func (h *UsersHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh-token")
	if err != nil {
		zap.S().Error(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "refresh()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token, err := h.usersService.Refresh(context.TODO(), cookie.Value)
	if err != nil {
		zap.S().Error(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "refresh()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := json.Marshal(map[string]string{
		"token": token.AccessToken,
	})
	if err != nil {
		zap.S().Error(
			zap.String("package", "internal/transport/rest"),
			zap.String("file", "user.go"),
			zap.String("function", "refresh()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", token.RefreshToken))
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
