package errors

import "errors"

var (
	ErrNotFound                     = errors.New("not found")
	ErrNameRequired                 = errors.New("name must not be empty")
	ErrNameTooShort                 = errors.New("name must be at least 3 characters")
	ErrInvalidEmail                 = errors.New("invalid e-mail format")
	ErrInvalidCPF                   = errors.New("invalid cpf format")
	ErrClienteAlreadyExistsForID    = errors.New("cliente with the provided id already exists")
	ErrClienteAlreadyExistsForCPF   = errors.New("cliente with the provided cpf already exists")
	ErrClienteAlreadyExistsForEmail = errors.New("cliente with the provided email already exists")
)
