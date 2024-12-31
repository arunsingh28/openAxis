package project

import (
	"encoding/json"
	"net/http"
)

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func InitProject(w http.ResponseWriter, r *http.Request) {
	project := Project{
		ID:   1,
		Name: "Project 1",
	}

	json.NewEncoder(w).Encode(project)
}
