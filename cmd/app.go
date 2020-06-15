package main

import (
	"net/http"
	"phraseservice/cmd/internal/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	srv := http.Server{
		Addr:    ":8080",
		Handler: getRouter(),
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
	// todo: graceful shutdown
}

func getRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(
		middleware.Recoverer,
		middleware.Logger,
		middleware.AllowContentType("application/json"),
	)

	router.Post("/get-phrase-hash", service.PhraseHashHandler)

	return router
}
