package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/with-insomnia/profile/internal/model"
	"github.com/with-insomnia/profile/internal/repository"
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

	username, ok := r.Context().Value(keyUser).(string)
	if !ok {
		fmt.Println(1)
		ErrorHandler(w, http.StatusUnauthorized)
		return
	}

	project.Author = username
	project.Created = time.Now()
	project.Updated = time.Now()
	err = h.repo.InsertNewProject(project)
	if err != nil {
		if err == repository.ErrNoUser {
			fmt.Println(err)
			fmt.Println(2)
			ErrorHandler(w, http.StatusUnauthorized)
			return
		}
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
