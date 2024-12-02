package ports

import (
	"github.com/filipeandrade6/fiap-pedeai-clientes/domain/entities"
)

type Repository interface {
	Create(cliente entities.Cliente) error
	List() ([]*entities.Cliente, error)
	GetClienteById(id entities.ID) (*entities.Cliente, error)
	GetClienteByCPF(cpf string) (*entities.Cliente, error)
	GetClienteByEmail(email string) (*entities.Cliente, error)
	Update(cliente entities.Cliente) error
	Remove(id entities.ID) error
}
