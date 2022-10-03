package server

import "errors"

var (
	ErrJSONUnmarshal = errors.New("couldn't unmarshal model into JSON")
	ErrJSONMarshal   = errors.New("couldn't marshal JSON into model")
	ErrDBCreate      = errors.New("couldn't create DB connection")
)
