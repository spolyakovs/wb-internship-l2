package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/spolyakovs/wb-internship-l2/develop/dev11/internal/model"
)

type eventRepository struct {
	store *Store
}

func (eRep *eventRepository) Create(ctx context.Context, event *model.Event) error {
	if !event.Validate() {
		return fmt.Errorf("couldn't create event: %+v\n\t%w", event, ErrValidation)
	}

	createQuery := "INSERT INTO events " +
		"(name, description, date) " +
		"VALUES ($1, $2, $3) RETURNING id;"

	if err := eRep.store.db.QueryRowContext(ctx, createQuery,
		event.Name,
		event.Description,
		event.DateString,
	).Scan(&event.ID); err != nil {
		err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
		return fmt.Errorf("couldn't create event: %+v\n\t%w", event, err)
	}

	return nil
}

func (eRep *eventRepository) Update(ctx context.Context, event *model.Event) error {
	if !event.Validate() {
		return fmt.Errorf("couldn't update event: %+v\n\t%w", event, ErrValidation)
	}

	if event.ID <= 0 {
		return fmt.Errorf("couldn't update event: %+v\n\t%w", event, ErrValidation)
	}

	updateQuery := "UPDATE events " +
		"SET (name, description, date) = " +
		"($2, $3, $4) WHERE id = $1;"

	updateRes, err := eRep.store.db.ExecContext(ctx, updateQuery,
		event.ID,
		event.Name,
		event.Description,
		event.DateString,
	)
	if err != nil {
		err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
		return fmt.Errorf("couldn't update event: %+v\n\t%w", event, err)
	}
	rows, err := updateRes.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrSQLInternal, err)
	}
	if rows != 1 {
		err = fmt.Errorf("expected to affect 1 row, affected %d", rows)
		return fmt.Errorf("%w: %v", ErrSQLInternal, err)
	}

	return nil
}

func (eRep *eventRepository) Delete(ctx context.Context, id uint64) error {
	deleteQuery := "DELETE FROM events " +
		"WHERE id = $1;"

	deleteRes, err := eRep.store.db.ExecContext(ctx, deleteQuery,
		id,
	)
	if err != nil {
		err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
		return fmt.Errorf("couldn't delete event with id: %v\n\t%w", id, err)
	}
	rows, err := deleteRes.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrSQLInternal, err)
	}
	if rows > 1 {
		err = fmt.Errorf("expected to affect max 1 row, affected %d", rows)
		return fmt.Errorf("%w: %v", ErrSQLInternal, err)
	}

	return nil
}

func (eRep *eventRepository) FindByID(ctx context.Context, id uint64) (*model.Event, error) {
	event := model.Event{}
	findByIdQuery := "SELECT " +
		"id, name, description, TO_CHAR(date, 'yyyy-mm-dd') " +
		"FROM events WHERE id = $1 LIMIT 1;"
	var dateString string

	if err := eRep.store.db.QueryRowContext(ctx,
		findByIdQuery,
		id,
	).Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&dateString,
	); err != nil {
		if err != sql.ErrNoRows {
			err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
		} else {
			err = fmt.Errorf("%w: %v", ErrNotExist, err)
		}
		return nil, fmt.Errorf("couldn't find event with id: %v\n\t%w", id, err)
	}

	if err := event.SetDate(dateString); err != nil {
		err = fmt.Errorf("couldn't set date (%v): %w", dateString, err)
		return nil, fmt.Errorf("couldn't find event with id: %v\n\t%w", id, err)
	}

	return &event, nil
}

func (eRep *eventRepository) FindAllByDatePart(ctx context.Context, timePart string) ([]*model.Event, error) {
	switch timePart {
	case "day", "week", "month":
		// do nothing
	default:
		return nil, fmt.Errorf("couldn't find items: wrong timePart %v", timePart)
	}

	events := make([]*model.Event, 0)
	findBetweenQuery := "SELECT " +
		"id, name, description, TO_CHAR(date, 'yyyy-mm-dd') " +
		"FROM events WHERE date_trunc($1, date) = date_trunc($1, now());"
	var dateString string

	rows, err := eRep.store.db.QueryContext(ctx,
		findBetweenQuery,
		timePart,
	)
	if err != nil {
		err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
		return nil, fmt.Errorf("couldn't find items for current %v\n\t%w", timePart, err)
	}
	defer rows.Close()

	for rows.Next() {
		event := model.Event{}
		if err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&dateString,
		); err != nil {
			err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
			return nil, fmt.Errorf("couldn't scan event into model\n\t%w", err)
		}

		if err := event.SetDate(dateString); err != nil {
			err = fmt.Errorf("couldn't set date (%v): %w", dateString, err)
			return nil, fmt.Errorf("couldn't scan event into model\n\t%w", err)
		}
		events = append(events, &event)
	}

	return events, nil
}
