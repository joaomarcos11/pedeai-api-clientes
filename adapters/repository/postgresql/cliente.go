package postgresql

import (
	"context"
	"errors"
	"fmt"

	"github.com/filipeandrade6/fiap-pedeai-clientes/adapters/repository/postgresql/db"
	"github.com/filipeandrade6/fiap-pedeai-clientes/domain/entities"
	entityErr "github.com/filipeandrade6/fiap-pedeai-clientes/domain/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repository) Create(cliente entities.Cliente) error {
	_, err := r.db.CreateCliente(
		context.Background(),
		db.CreateClienteParams{
			ID:    pgtype.UUID{Bytes: cliente.Id(), Valid: true},
			Nome:  pgtype.Text{String: cliente.Name(), Valid: true},
			Cpf:   pgtype.Text{String: cliente.CPF(), Valid: true},
			Email: pgtype.Text{String: cliente.Email(), Valid: true},
			Ativo: cliente.Active(),
		},
	)
	if err != nil {
		return fmt.Errorf("db creating cliente: %w", err)
	}

	return nil
}

func (r *Repository) List() ([]*entities.Cliente, error) {
	clientes, err := r.db.ListCliente(context.Background())
	if err != nil {
		return nil, err
	}

	var clientesOut []*entities.Cliente
	for _, cliente := range clientes {
		c, err := entities.New(
			cliente.ID.Bytes,
			cliente.Nome.String,
			cliente.Cpf.String,
			cliente.Email.String,
			cliente.Ativo,
		)
		if err != nil {
			return nil, err
		}

		clientesOut = append(clientesOut, c)
	}

	return clientesOut, nil

}

func (r *Repository) GetClienteById(id entities.ID) (*entities.Cliente, error) {
	c, err := r.db.GetClienteById(context.Background(), pgtype.UUID{Bytes: id, Valid: true})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entityErr.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	cOut, err := entities.New(
		c.ID.Bytes,
		c.Nome.String,
		c.Cpf.String,
		c.Email.String,
		c.Ativo,
	)
	if err != nil {
		return nil, err
	}

	return cOut, nil
}

func (r *Repository) GetClienteByCPF(cpf string) (*entities.Cliente, error) {
	c, err := r.db.GetClienteByCPF(context.Background(), pgtype.Text{String: cpf, Valid: true})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entityErr.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	cOut, err := entities.New(
		c.ID.Bytes,
		c.Nome.String,
		c.Cpf.String,
		c.Email.String,
		c.Ativo,
	)
	if err != nil {
		return nil, err
	}

	return cOut, nil
}

func (r *Repository) GetClienteByEmail(email string) (*entities.Cliente, error) {
	c, err := r.db.GetClienteByEmail(context.Background(), pgtype.Text{String: email, Valid: true})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entityErr.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	cOut, err := entities.New(
		c.ID.Bytes,
		c.Nome.String,
		c.Cpf.String,
		c.Email.String,
		c.Ativo,
	)
	if err != nil {
		return nil, err
	}

	return cOut, nil
}

func (r *Repository) Update(cliente entities.Cliente) error {
	err := r.db.UpdateCliente(context.Background(), db.UpdateClienteParams{
		ID:    pgtype.UUID{Bytes: cliente.Id(), Valid: true},
		Nome:  pgtype.Text{String: cliente.Name(), Valid: true},
		Cpf:   pgtype.Text{String: cliente.CPF(), Valid: true},
		Email: pgtype.Text{String: cliente.Email(), Valid: true},
		Ativo: cliente.Active(),
	})
	if err != nil {
		return fmt.Errorf("updating cliente %s in dabatabse: %w", cliente.Id(), err)
	}

	return nil
}

func (r *Repository) Remove(id entities.ID) error {
	err := r.db.DeleteCliente(context.Background(), pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return fmt.Errorf("removing cliente %s in database: %w", id, err)
	}

	return nil
}
