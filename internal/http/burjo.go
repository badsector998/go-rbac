package http

import (
	"encoding/json"
	"net/http"

	"github.com/badsector998/go-rbac/internal/domain"
)

type BurjoCreateUpdate struct {
	Name         string `json:"name"`
	Address      string `json:"address"`
	BurjoGroupID uint   `json:"burjo_group_id"`
}

func (s *Server) BurjoCreate(w http.ResponseWriter, r *http.Request) {
	var b BurjoCreateUpdate
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	burjo := domain.Burjo{
		Name:         b.Name,
		Address:      b.Address,
		BurjoGroupID: b.BurjoGroupID,
	}

	if err := s.app.Commands.BurjoCreate.Handle(r.Context(), burjo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
