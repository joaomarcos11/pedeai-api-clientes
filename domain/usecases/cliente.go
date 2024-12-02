package usecases

import (
	"github.com/google/uuid"

	"github.com/filipeandrade6/fiap-pedeai-clientes/domain/entities"
)

type ClienteUseCase interface {
	Create(cliente entities.Cliente) (uuid.UUID, error)
	List() ([]*entities.Cliente, error)
	GetClienteById(id uuid.UUID) (*entities.Cliente, error)
	GetClienteByCPF(cpf string) (*entities.Cliente, error)
	GetClienteByEmail(email string) (*entities.Cliente, error)
	Update(cliente entities.Cliente) error
	Remove(id uuid.UUID) error
}
