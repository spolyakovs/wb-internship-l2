package server

import "context"

// TODO
func (srv *Server) configureRouter(ctx context.Context) {
	srv.router.HandleFunc("/create_event", srv.handleCreateEvent(ctx))
	srv.router.HandleFunc("/update_event", srv.handleUpdateEvent(ctx))
	srv.router.HandleFunc("/delete_event", srv.handleDeleteEvent(ctx))

	srv.router.HandleFunc("/events_for_day", srv.handleEventsForDay(ctx))
	srv.router.HandleFunc("/events_for_week", srv.handleEventsForWeek(ctx))
	srv.router.HandleFunc("/events_for_month", srv.handleEventsForMonth(ctx))
}
