package http

import (
	"net/http"

	"github.com/badsector998/go-rbac/internal/app"
	"github.com/badsector998/go-rbac/internal/domain"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	app *app.App
}

func NewServer(a *app.App) *Server {
	return &Server{app: a}
}

func (s *Server) Run() error {
	r := chi.NewRouter()

	r.Get("burjo-group", s.BurjoGroupCreate)

	return http.ListenAndServe(":3001", r)
}

func (s *Server) BurjoGroupCreate(w http.ResponseWriter, r *http.Request) {
	s.app.Commands.BurjoGroupCreate.Handle(r.Context(), domain.BurjoGroup{})
}
