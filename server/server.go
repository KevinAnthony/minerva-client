package server

import (
	native "net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/kevinanthony/gorps/v2/http"
)

type HTTPServer struct {
	mux *chi.Mux
}

func NewServer(
	reqh http.RequestHandler,
) HTTPServer {
	server := HTTPServer{
		mux: chi.NewRouter(),
	}
	server.mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	return server
}

func (s HTTPServer) Run() {
	svr := &native.Server{
		Addr:              ":8080",
		Handler:           s.mux,
		ReadHeaderTimeout: 5 * time.Second, //TODO make configurable
	}

	svr.SetKeepAlivesEnabled(false)

	err := svr.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
