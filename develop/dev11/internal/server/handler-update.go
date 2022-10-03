package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/spolyakovs/wb-internship-l2/develop/dev11/internal/model"
	"github.com/spolyakovs/wb-internship-l2/develop/dev11/internal/store"
)

func (srv *Server) handleUpdateEvent(ctx context.Context) http.HandlerFunc {
	type request struct {
		ID          uint64 `json:"id"`
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
		event.ID = requestStruct.ID

		if _, err := srv.store.Events().FindByID(ctx, event.ID); err != nil {
			if err != nil {
				if errors.Is(err, store.ErrNotExist) {
					srv.error(w, req, http.StatusBadRequest, err)
					return
				}
				srv.error(w, req, http.StatusInternalServerError, err)
				return
			}
		}

		if err := srv.store.Events().Update(ctx, event); err != nil {
			if err != nil {
				srv.error(w, req, http.StatusInternalServerError, err)
				return
			}
		}

		srv.respond(w, req, http.StatusOK, &event)
	}
}
