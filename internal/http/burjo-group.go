package http

import (
	"encoding/json"
	"net/http"

	"github.com/badsector998/go-rbac/internal/domain"
)

type BurjoGroupCreateUpdate struct {
	Name        string `json:"name"`
	Owner       string `json:"owner"`
	Description string `json:"description"`
}

func (s *Server) BurjoGroupCreate(w http.ResponseWriter, r *http.Request) {
	var b BurjoGroupCreateUpdate

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bg := domain.BurjoGroup{
		Name:        b.Name,
		Owner:       b.Owner,
		Description: b.Description,
	}

	err = s.app.Commands.BurjoGroupCreate.Handle(r.Context(), bg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) BurjoGroupUpdate(w http.ResponseWriter, r *http.Request) {
	var b BurjoGroupCreateUpdate

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bg := domain.BurjoGroup{
		Name:        b.Name,
		Owner:       b.Owner,
		Description: b.Description,
	}

	err = s.app.Commands.BurjoGroupUpdate.Handle(r.Context(), bg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
