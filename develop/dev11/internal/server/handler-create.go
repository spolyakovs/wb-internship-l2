package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/spolyakovs/wb-internship-l2/develop/dev11/internal/model"
)

func (srv *Server) handleCreateEvent(ctx context.Context) http.HandlerFunc {
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Date        string `json:"date"`
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

		event, err := model.NewEvent(
			requestStruct.Name,
			requestStruct.Description,
			requestStruct.Date,
		)
		if err != nil {
			err = fmt.Errorf("%w: %v", ErrJSONMarshal, err)
			srv.error(w, req, http.StatusBadRequest, err)
			return
		}

		if err := srv.store.Events().Create(ctx, event); err != nil {
			if err != nil {
				srv.error(w, req, http.StatusInternalServerError, err)
				return
			}
		}

		srv.respond(w, req, http.StatusOK, &event)
	}
}
