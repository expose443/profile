package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/with-insomnia/profile/internal/model"
)

func (h *Handlers) CreateProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	var project model.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	project.Created = time.Now()
	project.Updated = time.Now()
	err = h.repo.InsertNewProject(project)
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func (h *Handlers) GetProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	projects, err := h.repo.GetAllProjects()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(projects)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	return
}
