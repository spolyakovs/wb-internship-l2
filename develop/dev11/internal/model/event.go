package model

import "time"

type Event struct {
	ID          uint64    `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"-"`
	DateString  string    `json:"date"` // in format (2006-01-02) (ISO 8601)
}

func NewEvent(name string, description string, dateString string) (*Event, error) {
	event := &Event{
		Name:        name,
		Description: description,
	}
	if err := event.SetDate(dateString); err != nil {
		return nil, err
	}
	return event, nil
}

func (e *Event) SetDate(dateString string) error {
	e.DateString = dateString
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return err
	}

	e.Date = date

	return nil
}

func (e *Event) Validate() bool {
	return e.Name != ""
}
