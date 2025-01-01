package project

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/arunsingh28/openaxis/internal/types"
	response "github.com/arunsingh28/openaxis/internal/utils"
	"github.com/go-playground/validator/v10"
)

func InitProject(w http.ResponseWriter, r *http.Request) {
	var project types.Project

	err := json.NewDecoder(r.Body).Decode(&project)

	if errors.Is(err, io.EOF) {
		response.WriteError(w, http.StatusBadRequest, errors.New("request body is empty"))
		return
	}

	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := validator.New().Struct(project); err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))
		return
	}

	slog.Info("project initialized", slog.String("name", project.Name))

	response.WriteJson(w, http.StatusCreated, map[string]string{"message": "project initialized"})
}
