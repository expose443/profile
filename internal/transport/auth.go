package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/with-insomnia/profile/internal/model"
	"github.com/with-insomnia/profile/internal/repository"
)

func NewHandler(repo repository.Repository) Handlers {
	return Handlers{
		repo: repo,
	}
}

type Handlers struct {
	repo repository.Repository
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var credintails model.Credintails
	err := json.NewDecoder(r.Body).Decode(&credintails)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(credintails)
}
