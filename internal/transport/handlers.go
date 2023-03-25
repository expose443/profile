package transport

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/with-insomnia/profile/internal/model"
)

func NewHandler(db *sql.DB) Handlers {
	return Handlers{
		db: db,
	}
}

type Handlers struct {
	db *sql.DB
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var credintails model.Credintails
	err := json.NewDecoder(r.Body).Decode(&credintails)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(credintails)
}
