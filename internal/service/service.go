package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cyaniccerulean.com/inventory/v2/internal/db"
	"cyaniccerulean.com/inventory/v2/internal/model"
	"cyaniccerulean.com/inventory/v2/internal/routes"
	_ "github.com/mattn/go-sqlite3"
)

type Service struct {
	config model.Config
	db     *db.Database
}

// set up new service
func New(ctx context.Context, config model.Config) (Service, error) {
	db, err := db.Initialize(ctx, config.DBPath)
	return Service{config, db}, err
}

func (s Service) Run(ctx context.Context) error {
	// register HTTP routes
	s.registerRoutes()

	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// start HTTP server
	srv := &http.Server{Addr: fmt.Sprintf(":%d", s.config.Port)}
	go func() {
		log.Printf("Server starting on :%d\n", s.config.Port)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Println(err)
			srv.Shutdown(newCtx)
			cancel()
		}
	}()

	// wait for app stop signal, clean up
	<-newCtx.Done()
	log.Println("stopping server")
	srv.Shutdown(newCtx)
	s.db.Close()

	return nil
}

func (s Service) registerRoutes() {
	api := routes.InitAPI(s.db)
	pages := routes.InitPages(s.config)

	// register the index page handler
	http.HandleFunc("/", pages.IndexHandler)

	// static assets
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/static/assets"))))

	// health endpoint
	http.HandleFunc("/health", routes.HandleHealthEndpoint)

	// register the API handler functions
	http.HandleFunc("/api/entry", api.APIHandler)

	// register the entry page handler
	http.HandleFunc("/entry", pages.EntryHandler)
}
