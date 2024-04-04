package app

import (
	"net/http"

	"JWT/pkg/middleware"
	"JWT/pkg/service"
	"JWT/pkg/store"
)

type Implementation struct {
	services *service.Service
	storage  store.Store
	midware  *middleware.Middleware
}

func NewImplementation(srv *service.Service, storage store.Store, midware *middleware.Middleware) *Implementation {
	return &Implementation{
		services: srv,
		storage:  storage,
		midware:  midware,
	}
}

func (i *Implementation) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", i.Registration)
	mux.HandleFunc("/sign-in", i.Login)

	final := http.HandlerFunc(i.Interact)

	mux.Handle("/api/interact", i.midware.UserIdentity(final))

	return mux
}
