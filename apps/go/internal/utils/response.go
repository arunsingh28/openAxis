package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	if status == 0 {
		status = http.StatusOK
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, Response{Status: status, Error: err.Error()})
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errors []string
	for _, e := range errs {
		switch e.ActualTag() {
		case "required":
			errors = append(errors, fmt.Sprintf("filed %s is required", e.Field()))
		default:
			errors = append(errors, fmt.Sprintf("field %s is invalid", e.Field()))
		}
	}
	return Response{
		Status: http.StatusBadRequest, Error: strings.Join(errors, ", "),
	}
}
