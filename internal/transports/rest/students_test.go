package rest

// import (
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"io"
// 	"net/http"
// 	"strconv"
// 	"test-crud/internal/model"
// 	"test-crud/pkg/logger"

// 	"github.com/gorilla/mux"
// 	"go.uber.org/zap"
// )

// // @Summary Create student
// // @Description Create student via http request
// // @Tags students
// // @Accept json
// // @Produce json
// // @Param student body model.CreateStudentInput true "Data for creating student"
// // @Success 202 {string} string "Accepted"
// // @Failure 400 {string} string "Bad request"
// // @Failure 500 {string} string "Internal Server Error"
// // @Router /students [post]
// func (h *Handler) createStudent(w http.ResponseWriter, r *http.Request) {
// 	reqBytes, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		logger.Fatal(
// 			zap.String("package", "transport/rest"),
// 			zap.String("file", "students.go"),
// 			zap.String("function", "createStudent()"),
// 			zap.Error(err),
// 		)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	var student model.CreateStudentInput

// 	if err = json.Unmarshal(reqBytes, &student); err != nil {
// 		logger.Fatal(
// 			zap.String("package", "transport/rest"),
// 			zap.String("file", "students.go"),
// 			zap.String("function", "createStudent()"),
// 			zap.Error(err),
// 		)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	if err := h.services.Students.Create(context.TODO(), student); err != nil {
// 		logger.Fatal(
// 			zap.String("package", "transport/rest"),
// 			zap.String("file", "students.go"),
// 			zap.String("function", "createStudent()"),
// 			zap.Error(err),
// 		)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusAccepted)
// }

// // @Summary Delete student
// // @Description Delete student via http request
// // @Tags students
// // @Accept json
// // @Produce json
// // @Param id path int true "ID for deleting student"
// // @Success 200 {string} string "OK"
// // @Failure 400 {string} string "Bad request"
// // @Failure 500 {string} string "Internal Server Error"
// // @Router /students/{id} [delete]
// func (h *Handler) deleteStudent(w http.ResponseWriter, r *http.Request) {
// 	id, err := getIdFromRequest(r)
// 	if err != nil {
// 		logger.Fatal(
// 			zap.String("package", "transport/rest"),
// 			zap.String("file", "students.go"),
// 			zap.String("function", "deleteStudent()"),
// 			zap.Error(err),
// 		)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	if err := h.services.Students.Delete(context.TODO(), id); err != nil {
// 		logger.Fatal(
// 			zap.String("package", "transport/rest"),
// 			zap.String("file", "students.go"),
// 			zap.String("function", "deleteStudent()"),
// 			zap.Error(err),
// 		)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// }

// // @Summary Get students
// // @Description Get all students via http request
// // @Tags students
// // @Accept json
// // @Produce json
// // @Success 200 {array} model.Student "Accepted"
// // @Failure 400 {string} string "Bad request"
// // @Failure 500 {string} string "Internal Server Error"
// // @Router /students [get]
// func (h *Handler) getAllStudents(w http.ResponseWriter, r *http.Request) {
// 	students, err := h.services.Students.GetAll(context.TODO())
// 	if err != nil {
// 		logger.Fatal(
// 			zap.String("package", "transport/rest"),
// 			zap.String("file", "students.go"),
// 			zap.String("function", "getAllStudents()"),
// 			zap.Error(err),
// 		)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	response, err := json.Marshal(students)
// 	if err != nil {
// 		logger.Fatal(
// 			zap.String("package", "transport/rest"),
// 			zap.String("file", "students.go"),
// 			zap.String("function", "getAllStudents()"),
// 			zap.Error(err),
// 		)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Add("Context-Type", "application/json")
// 	w.Write(response)
// }

// // @Summary Get student
// // @Description Get student by id via http request
// // @Tags students
// // @Accept json
// // @Produce json
// // @Param id path int true "ID for getting student"
// // @Success 200 {object} model.Student "Accepted"
// // @Failure 400 {string} string "Bad request"
// // @Failure 500 {string} string "Internal Server Error"
// // @Router /students/{id} [get]
// func (h *Handler) getStudentById(w http.ResponseWriter, r *http.Request) {
// 	id, err := getIdFromRequest(r)
// 	if err != nil {
// 		logger.Fatal(
// 			zap.String("package", "transport/rest"),
// 			zap.String("file", "students.go"),
// 			zap.String("function", "getStudentById()"),
// 			zap.Error(err),
// 		)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	student, err := h.services.Students.GetById(context.TODO(), id)
// 	if err != nil {
// 		logger.Fatal(
// 			zap.String("package", "transport/rest"),
// 			zap.String("file", "students.go"),
// 			zap.String("function", "getStudentById()"),
// 			zap.Error(err),
// 		)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	response, err := json.Marshal(student)
// 	if err != nil {
// 		logger.Fatal(
// 			zap.String("package", "transport/rest"),
// 			zap.String("file", "students.go"),
// 			zap.String("function", "getStudentById()"),
// 			zap.Error(err),
// 		)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Add("Context-Type", "application/json")
// 	w.Write(response)
// }

// // @Summary Update student
// // @Description Update student via HTTP request
// // @Tags students
// // @Accept json
// // @Produce json
// // @Param id path int true "ID for updating student"
// // @Param request body model.UpdateStudentInput true "New information for update"
// // @Success 200 {string} string "OK"
// // @Failure 400 {string} string "Bad request"
// // @Failure 500 {string} string "Internal Server Error"
// // @Router /students/{id} [patch]
// func (h *Handler) updateStudent(w http.ResponseWriter, r *http.Request) {
// 	id, err := getIdFromRequest(r)
// 	if err != nil {
// 		logger.Fatal(
// 			zap.String("package", "transport/rest"),
// 			zap.String("file", "students.go"),
// 			zap.String("function", "updateStudent()"),
// 			zap.Error(err),
// 		)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	reqBytes, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		logger.Fatal(
// 			zap.String("package", "transport/rest"),
// 			zap.String("file", "students.go"),
// 			zap.String("function", "updateStudent()"),
// 			zap.Error(err),
// 		)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	var student model.UpdateStudentInput
// 	if err := json.Unmarshal(reqBytes, &student); err != nil {
// 		logger.Fatal(
// 			zap.String("package", "transport/rest"),
// 			zap.String("file", "students.go"),
// 			zap.String("function", "updateStudent()"),
// 			zap.Error(err),
// 		)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	err = h.services.Students.Update(context.TODO(), id, student)
// 	if err != nil {
// 		logger.Fatal(
// 			zap.String("package", "transport/rest"),
// 			zap.String("file", "students.go"),
// 			zap.String("function", "updateStudent()"),
// 			zap.Error(err),
// 		)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// }

// func getIdFromRequest(r *http.Request) (int64, error) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.ParseInt(vars["id"], 10, 64)
// 	if err != nil {
// 		return 0, err
// 	}
// 	if id == 0 {
// 		return 0, errors.New("Id couldn't be zero")
// 	}
// 	return id, nil
// }
