package server

import (
	"context"
	"errors"
	"net/http"
)

func (srv *Server) handleEventsForWeek(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			srv.error(w, req, http.StatusMethodNotAllowed, errors.New("wrong request method"))
			return
		}

		events, err := srv.store.Events().FindAllByDatePart(ctx, "week")
		if err != nil {
			if err != nil {
				srv.error(w, req, http.StatusInternalServerError, err)
				return
			}
		}

		srv.respond(w, req, http.StatusOK, events)
	}
}
