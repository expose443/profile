package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/with-insomnia/profile/internal/model"
	"github.com/with-insomnia/profile/internal/repository"
)

func NewHandler(repo *repository.Repository) *Handlers {
	return &Handlers{
		repo: *repo,
	}
}

type Handlers struct {
	repo repository.Repository
}

func (h *Handlers) Status(w http.ResponseWriter, r *http.Request) {
	status := make(map[string]bool)
	status["status - ok"] = true
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&status)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	var credintails model.Credintails
	err := json.NewDecoder(r.Body).Decode(&credintails)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		// fmt.Println(err)
		return
	}
	// fmt.Println(credintails)
	err = h.repo.CheckUser(credintails)
	if err != nil {
		ErrorHandler(w, http.StatusUnauthorized)
		// fmt.Println(err)
		return
	}
	token, err := GenerateJWT(credintails.Username)
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	fmt.Println(token)
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt_token",
		Value:    token,
		HttpOnly: true,
	})
	w.WriteHeader(200)
	return
}
