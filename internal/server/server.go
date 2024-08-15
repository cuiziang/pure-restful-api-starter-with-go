package server

import (
	"database/sql"
	"fmt"
	"github.com/cuiziang/pure-restFul-api-starter-with-go/internal/handlers"
	"github.com/cuiziang/pure-restFul-api-starter-with-go/internal/log"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

type HandlerFunc func(http.Handler) http.Handler

type HandlerCreator func(*sql.DB) http.Handler

type Route struct {
	Path        string
	Handler     http.Handler
	Middlewares []HandlerFunc
}

type Server struct {
	mux    *http.ServeMux
	db     *sql.DB
	routes []Route
}

func NewServer() *Server {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return &Server{
		db:     db,
		mux:    http.NewServeMux(),
		routes: make([]Route, 0),
	}
}

func (s *Server) AddRoute(path string, creator HandlerCreator, middlewares ...HandlerFunc) {
	handler := creator(s.db)
	s.routes = append(s.routes, Route{
		Path:        path,
		Handler:     handler,
		Middlewares: middlewares,
	})
}

func (s *Server) SetupRoutes() *Server {
	s.AddRoute("/health", handlers.NewHealthHandler, log.LoggingMiddleware)
	s.AddRoute("/", handlers.NewHomeHandler, log.LoggingMiddleware)

	for _, route := range s.routes {
		handler := route.Handler
		for i := len(route.Middlewares) - 1; i >= 0; i-- {
			handler = route.Middlewares[i](handler)
		}
		s.mux.Handle(route.Path, handler)
	}
	return s
}

func (s *Server) Start() error {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	err := http.ListenAndServe(port, s.mux)
	if err != nil {
		return err
	}
	fmt.Sprintf("Server started on port %s", port)
	return nil
}
