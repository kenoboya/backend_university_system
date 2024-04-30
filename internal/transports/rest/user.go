package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"test-crud/internal/model"
	"test-crud/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type UsersHandler struct {
	usersService  service.Users
	peopleService service.People
}

func NewUsersHandler(usersService service.Users, peopleService service.People) *UsersHandler {
	return &UsersHandler{
		usersService:  usersService,
		peopleService: peopleService,
	}
}

func (h *UsersHandler) initRoutes(router *mux.Router) {
	users := router.PathPrefix("/users").Subrouter()
	{
		users.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)
		users.HandleFunc("/sign-in", h.signIn).Methods(http.MethodPost)

		people := router.PathPrefix("/people").Subrouter()
		{
			people.HandleFunc("", h.createPerson).Methods(http.MethodPost)
			people.HandleFunc("", h.getPeople).Methods(http.MethodGet)
			people.HandleFunc("/{id:[0-9]+}", h.getPerson).Methods(http.MethodPost)
			people.HandleFunc("/{id:[0-9]+}", h.updatePerson).Methods(http.MethodPatch)
			people.HandleFunc("/{id:[0-9]+}", h.deletePerson).Methods(http.MethodDelete)
		}
	}
}

// @Summary User registration
// @Description User registration
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.UserSignUpInput true "Data for registration user"
// @Success 201 {string} string "Registered"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/sign-up [post]
func (h *UsersHandler) signUp(w http.ResponseWriter, r *http.Request) {
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
func (h *UsersHandler) signIn(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
