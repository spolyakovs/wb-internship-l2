package store

import "database/sql"

type Store struct {
	db              *sql.DB
	eventRepository *eventRepository
}

func New(db *sql.DB) *Store {
	newStore := &Store{
		db: db,
	}

	return newStore
}

func (st *Store) Events() *eventRepository {
	if st.eventRepository == nil {
		st.eventRepository = &eventRepository{
			store: st,
		}
	}

	return st.eventRepository
}
