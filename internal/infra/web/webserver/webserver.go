package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Webserver struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *Webserver {
	return &Webserver{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *Webserver) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *Webserver) Start() {
	//s.Router.Use((middleware.Logger))
	for path, handler := range s.Handlers {
		s.Router.Get(path, handler)
		s.Router.Post(path, handler)
	}

	http.ListenAndServe(s.WebServerPort, s.Router)
}
