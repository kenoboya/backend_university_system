package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"test-crud/internal/model"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (h *Handler) initAdminEmployeesRoutes(admin *mux.Router) {
	employees := admin.PathPrefix("/employees").Subrouter()
	{
		employees.HandleFunc("", h.Admins.createEmployee).Methods(http.MethodPost)
		employees.HandleFunc("", h.Admins.getEmployees).Methods(http.MethodGet)
		employees.HandleFunc("/{id:[0-9]+}", h.Admins.getEmployee).Methods(http.MethodGet)
		employees.HandleFunc("/{id:[0-9]+}", h.Admins.updateEmployee).Methods(http.MethodPatch)
		employees.HandleFunc("/{id:[0-9]+}", h.Admins.deleteEmployee).Methods(http.MethodDelete)
	}
}

// @Summary create employee
// @Description create employee
// @Tags admin-employees
// @Accept json
// @Produce json
// @Param employee body model.CreateEmployeeInput true "Data for creating employee"
// @Success 202 {string} string "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/employees [post]
func (h *AdminsHandler) createEmployee(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_employee.go"),
			zap.String("function", "createEmployee()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var employee model.CreateEmployeeInput

	if err = json.Unmarshal(reqBytes, &employee); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_employee.go"),
			zap.String("function", "createEmployee()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Employees.Create(context.TODO(), employee); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_employee.go"),
			zap.String("function", "createEmployee()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

// @Summary Get employees
// @Description get employees
// @Tags admin-employees
// @Accept json
// @Produce json
// @Success 200 {array} model.Employee "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/employees [get]
func (h *AdminsHandler) getEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.services.Employees.GetAll(context.TODO())
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_employee.go"),
			zap.String("function", "getEmployees()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(employees)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_employee.go"),
			zap.String("function", "getEmployees()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Get employee
// @Description get employee by id
// @Tags admin-employees
// @Accept json
// @Produce json
// @Param id path int true "ID for getting employee"
// @Success 200 {object} model.Employee "Accepted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/employees/{id} [get]
func (h *AdminsHandler) getEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_employee.go"),
			zap.String("function", "getEmployee()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	employee, err := h.services.Employees.GetById(context.TODO(), id)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_employee.go"),
			zap.String("function", "getEmployee()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := json.Marshal(employee)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_employee.go"),
			zap.String("function", "getEmployee()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")
	w.Write(response)
}

// @Summary Update employee
// @Description update employee
// @Tags admin-employees
// @Accept json
// @Produce json
// @Param id path int true "ID for updating employee"
// @Param request body model.UpdateEmployeeInput true "New information for update"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/employees/{id} [patch]
func (h *AdminsHandler) updateEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_employee.go"),
			zap.String("function", "updateEmployee()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_employee.go"),
			zap.String("function", "updateEmployee()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var employee model.UpdateEmployeeInput
	if err := json.Unmarshal(reqBytes, &employee); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_employee.go"),
			zap.String("function", "updateEmployee()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.services.Employees.Update(context.TODO(), id, employee)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_employee.go"),
			zap.String("function", "updateEmployee()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary Delete employee
// @Description delete employee
// @Tags admin-employees
// @Accept json
// @Produce json
// @Param id path int true "ID for deleting employee"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /admin/hub/employees/{id} [delete]
func (h *AdminsHandler) deleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_employee.go"),
			zap.String("function", "deleteEmployee()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.Employees.Delete(context.TODO(), id); err != nil {
		zap.S().Error(
			zap.String("package", "transport/rest"),
			zap.String("file", "admin_employee.go"),
			zap.String("function", "deleteEmployee()"),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
