package repoerrors

import (
	"errors"
)

var (
	ErrNotInserted = errors.New("not inserted")
	ErrNotDeleted  = errors.New("not deleted")

	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")

	ErrNotEnoughBalance = errors.New("not enough balance")
)
