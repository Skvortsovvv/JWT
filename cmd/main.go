package main

import (
	"log"

	"JWT/pkg/app"
	"JWT/pkg/middleware"
	"JWT/pkg/model"
	"JWT/pkg/service"
	"JWT/pkg/store"
)

func main() {

	server := new(model.Server)
	db, err := store.NewPostgresDB()
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}
	storage := store.NewStorage(db)
	srv := service.NewService(storage)
	mdlw := middleware.NewMiddleware(storage, srv)
	impl := app.NewImplementation(srv, storage, mdlw)

	err = server.Run("8080", impl.InitRoutes())
	if err != nil {
		log.Fatalf("cant start server: %v", err)
	}

}
