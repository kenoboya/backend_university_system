package common

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIdFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, errors.New("id couldn't be zero")
	}
	return id, nil
}

func GetIdStringFromRequest(r *http.Request, nameID string) (string, error) {
	vars := mux.Vars(r)
	id := vars[nameID]
	if id == "" {
		return id, errors.New("id couldn't be empty")
	}
	return id, nil
}
