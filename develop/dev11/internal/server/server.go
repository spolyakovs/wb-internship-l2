package server

import (
	"context"
	"net/http"

	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/spolyakovs/wb-internship-l2/develop/dev11/internal/store"
)

type Server struct {
	router        *http.ServeMux
	logger        *logrus.Logger
	store         store.Store
	STANConnected bool
}

func NewServer(ctx context.Context, st store.Store, loggerLever logrus.Level) *Server {
	srv := &Server{
		router:        http.NewServeMux(),
		logger:        logrus.New(),
		store:         st,
		STANConnected: false,
	}

	srv.logger.Level = loggerLever

	srv.configureRouter(ctx)

	return srv
}

func (srv *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	srv.router.ServeHTTP(w, req)
}

func (srv *Server) error(w http.ResponseWriter, req *http.Request, code int, err error) {
	srv.logger.Debug(err)

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func (srv *Server) respond(w http.ResponseWriter, req *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"result": data})
	}
}
