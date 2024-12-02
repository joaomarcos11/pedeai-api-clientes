package entities

import (
	"net/mail"
	"regexp"
	"strings"

	entityErr "github.com/filipeandrade6/fiap-pedeai-clientes/domain/errors"
)

type Cliente struct {
	id     ID
	name   string
	cpf    string
	email  string
	active bool
}

var (
	cpfPattern = regexp.MustCompile(`^\d{11}$`)
)

func New(id ID, name, cpf, email string, active bool) (*Cliente, error) {
	c := Cliente{
		id:     id,
		name:   strings.TrimSpace(name),
		cpf:    cpf,
		email:  strings.TrimSpace(email),
		active: active,
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Cliente) Id() ID {
	return c.id
}

func (c *Cliente) Name() string {
	return c.name
}

func (c *Cliente) CPF() string {
	return c.cpf
}

func (c *Cliente) Email() string {
	return c.email
}

func (c *Cliente) Active() bool {
	return c.active
}

func (c *Cliente) Validate() error {
	if len(c.name) == 0 {
		return entityErr.ErrNameRequired
	}

	if len(c.name) <= 3 {
		return entityErr.ErrNameTooShort
	}

	if !(cpfPattern.MatchString(c.cpf)) {
		return entityErr.ErrInvalidCPF
	}

	if _, err := mail.ParseAddress(c.email); err != nil {
		return entityErr.ErrInvalidEmail
	}

	return nil
}
