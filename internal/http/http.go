package http

import (
	"net/http"

	"github.com/badsector998/go-rbac/internal/app"
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

	burjoGroupRouter := chi.NewRouter()
	burjoGroupRouter.Post("/create", s.BurjoGroupCreate)
	burjoGroupRouter.Put("/update", s.BurjoGroupUpdate)

	burjoRouter := chi.NewRouter()
	burjoRouter.Post("/create", s.BurjoCreate)

	r.Mount("/burjo-group", burjoGroupRouter)
	r.Mount("burjo", burjoRouter)

	return http.ListenAndServe(":3001", r)
}
