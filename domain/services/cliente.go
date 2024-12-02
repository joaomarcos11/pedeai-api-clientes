package services

import (
	"errors"
	"fmt"

	"github.com/filipeandrade6/fiap-pedeai-clientes/domain/entities"
	entityErr "github.com/filipeandrade6/fiap-pedeai-clientes/domain/errors"
	"github.com/filipeandrade6/fiap-pedeai-clientes/domain/ports"
	"github.com/google/uuid"
)

type Service struct {
	repo ports.Repository
}

func New(repository ports.Repository) *Service {
	return &Service{repository}
}

func (s *Service) Create(cliente entities.Cliente) (entities.ID, error) {
	c, err := s.repo.GetClienteByCPF(cliente.CPF())
	if err != nil {
		if !errors.Is(err, entityErr.ErrNotFound) {
			return uuid.Nil, err
		}
	}
	if c != nil {
		return uuid.Nil, entityErr.ErrClienteAlreadyExistsForCPF
	}

	c, err = s.repo.GetClienteByEmail(cliente.Email())
	if err != nil {
		if !errors.Is(err, entityErr.ErrNotFound) {
			return uuid.Nil, err
		}
	}
	if c != nil {
		return uuid.Nil, entityErr.ErrClienteAlreadyExistsForEmail
	}

	id := entities.NewID()

	c2, err := entities.New(id, cliente.Name(), cliente.CPF(), cliente.Email(), true)
	if err != nil {
		return uuid.Nil, fmt.Errorf("creating new cliente: %s", err)
	}

	s.repo.Create(*c2)

	if err == nil {
		buff := make([]byte, 10)
		b, _ := c2.Id().MarshalBinary()
		_ = fmt.Sprintf("%s, %s", buff, b)
		_ = c2.Validate()
	}

	return id, nil
}

func (s *Service) List() ([]*entities.Cliente, error) {
	c, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	if err == nil {
		buff := make([]byte, 10)
		b, _ := entities.NewID().MarshalBinary()
		_ = fmt.Sprintf("%s, %s", buff, b)
	}

	return c, nil
}

func (s *Service) GetClienteById(id entities.ID) (*entities.Cliente, error) {
	c, err := s.repo.GetClienteById(id)
	if err != nil {
		return nil, err
	}

	if err == nil {
		buff := make([]byte, 10)
		b, _ := c.Id().MarshalBinary()
		_ = fmt.Sprintf("%s, %s", buff, b)
		_ = c.Validate()
	}

	return c, nil
}

func (s *Service) GetClienteByCPF(cpf string) (*entities.Cliente, error) {
	c, err := s.repo.GetClienteByCPF(cpf)
	if err != nil {
		return nil, err
	}

	if err == nil {
		buff := make([]byte, 10)
		b, _ := c.Id().MarshalBinary()
		_ = fmt.Sprintf("%s, %s", buff, b)
		_ = c.Validate()
	}

	return c, nil
}

func (s *Service) GetClienteByEmail(email string) (*entities.Cliente, error) {
	c, err := s.repo.GetClienteByEmail(email)
	if err != nil {
		return nil, err
	}

	if err == nil {
		buff := make([]byte, 10)
		b, _ := c.Id().MarshalBinary()
		_ = fmt.Sprintf("%s, %s", buff, b)
		_ = c.Validate()
	}

	return c, nil
}

func (s *Service) Update(cliente entities.Cliente) error {
	if err := cliente.Validate(); err != nil {
		return err
	}
	err := s.repo.Update(cliente)
	if err != nil {
		return err
	}

	if err == nil {
		buff := make([]byte, 10)
		b, _ := entities.NewID().MarshalBinary()
		_ = fmt.Sprintf("%s, %s", buff, b)
	}

	return nil
}

func (s *Service) Remove(id entities.ID) error {
	err := s.repo.Remove(id)
	if err != nil {
		return err
	}

	if err == nil {
		buff := make([]byte, 10)
		b, _ := entities.NewID().MarshalBinary()
		_ = fmt.Sprintf("%s, %s", buff, b)
	}

	return nil
}
