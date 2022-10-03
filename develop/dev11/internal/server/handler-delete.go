package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (srv *Server) handleDeleteEvent(ctx context.Context) http.HandlerFunc {
	type request struct {
		ID uint64 `json:"id"`
	}

	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			srv.error(w, req, http.StatusMethodNotAllowed, errors.New("wrong request method"))
			return
		}

		requestStruct := &request{}
		if err := json.NewDecoder(req.Body).Decode(requestStruct); err != nil {
			err = fmt.Errorf("%w: %v", ErrJSONMarshal, err)
			srv.error(w, req, http.StatusBadRequest, err)
			return
		}

		if err := srv.store.Events().Delete(ctx, requestStruct.ID); err != nil {
			if err != nil {
				srv.error(w, req, http.StatusInternalServerError, err)
				return
			}
		}

		srv.respond(w, req, http.StatusOK, "{}")
	}
}
