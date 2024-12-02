package entities

import "github.com/filipeandrade6/fiap-pedeai-clientes/domain/entities"

type Cliente struct {
	ID     entities.ID `json:"id,omitempty"`
	Name   string      `json:"name,omitempty"`
	CPF    string      `json:"cpf,omitempty"`
	Email  string      `json:"email,omitempty"`
	Active bool        `json:"active,omitempty"`
}

func (c *Cliente) ToDomain() (*entities.Cliente, error) {
	cDomain, err := entities.New(c.ID, c.Name, c.CPF, c.Email, c.Active)
	if err != nil {
		return nil, err
	}

	return cDomain, nil
}

func FromDomain(c *entities.Cliente) (*Cliente, error) {
	return &Cliente{
		ID:     c.Id(),
		Name:   c.Name(),
		CPF:    c.CPF(),
		Email:  c.Email(),
		Active: c.Active(),
	}, nil
}
