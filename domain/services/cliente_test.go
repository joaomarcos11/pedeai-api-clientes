package services

import (
	"errors"
	"maps"
	"slices"
	"testing"

	"github.com/filipeandrade6/fiap-pedeai-clientes/domain/entities"
	entityErr "github.com/filipeandrade6/fiap-pedeai-clientes/domain/errors"
	"github.com/google/uuid"
)

var existentClientID string = "db6c3a54-541f-472c-8810-13508c930070"
var existentClientCPF string = "12312312312"
var existentClientEmail string = "fulano@email.com"

var clienteIdError string = "db6c3a54-541f-7777-8810-13508c930070"
var clienteCpfError string = "90867562301"
var clienteEmailError string = "error@email.com"
var flagListClienteError bool = false

type ClienteRepositoryMock struct {
	Base map[entities.ID]*entities.Cliente
}

var clienteRepoMock ClienteRepositoryMock

func init() {
	uuid, _ := entities.StringToID(existentClientID)

	c, _ := entities.New(
		uuid,
		"Fulano",
		existentClientCPF,
		existentClientEmail,
		true,
	)

	repoMock := ClienteRepositoryMock{}
	repoMock.Base = make(map[entities.ID]*entities.Cliente)
	repoMock.Base[uuid] = c

	clienteRepoMock = repoMock
}

func (c *ClienteRepositoryMock) Create(cliente entities.Cliente) error {
	if cliente.Id().String() == clienteIdError || cliente.CPF() == clienteCpfError || cliente.Email() == clienteEmailError {
		return errors.New("repo mock error")
	}

	for k, v := range c.Base {
		if k == cliente.Id() {
			return entityErr.ErrClienteAlreadyExistsForID
		}
		if v.CPF() == cliente.CPF() {
			return entityErr.ErrClienteAlreadyExistsForCPF
		}
		if v.Email() == cliente.Email() {
			return entityErr.ErrClienteAlreadyExistsForEmail
		}
	}

	c.Base[cliente.Id()] = &cliente
	return nil
}

func (c *ClienteRepositoryMock) List() ([]*entities.Cliente, error) {
	if flagListClienteError {
		return nil, errors.New("new mock error")
	}

	return slices.Collect(maps.Values(c.Base)), nil
}

func (c *ClienteRepositoryMock) GetClienteById(id entities.ID) (*entities.Cliente, error) {
	errCUuid, _ := entities.StringToID(clienteIdError)
	if id == errCUuid {
		return nil, errors.New("new mock error")
	}

	cliente, ok := c.Base[id]
	if !ok {
		return nil, entityErr.ErrNotFound
	}
	return cliente, nil
}

func (c *ClienteRepositoryMock) GetClienteByCPF(cpf string) (*entities.Cliente, error) {
	if cpf == clienteCpfError {
		return nil, errors.New("new mock error")
	}

	for _, cliente := range c.Base {
		if cliente.CPF() == cpf {
			return cliente, nil
		}
	}

	return nil, entityErr.ErrNotFound
}

func (c *ClienteRepositoryMock) GetClienteByEmail(email string) (*entities.Cliente, error) {
	if email == clienteEmailError {
		return nil, errors.New("new mock error")
	}

	for _, cliente := range c.Base {
		if cliente.Email() == email {
			return cliente, nil
		}
	}

	return nil, entityErr.ErrNotFound
}

func (c *ClienteRepositoryMock) Update(cliente entities.Cliente) error {
	errCUuid, _ := entities.StringToID(clienteIdError)
	if cliente.Id() == errCUuid {
		return errors.New("new mock error")
	}

	if _, ok := c.Base[cliente.Id()]; !ok {
		return entityErr.ErrNotFound
	}

	c.Base[cliente.Id()] = &cliente

	return nil
}

func (c *ClienteRepositoryMock) Remove(id entities.ID) error {
	errCUuid, _ := entities.StringToID(clienteIdError)
	if id == errCUuid {
		return errors.New("new mock error")
	}

	if _, ok := c.Base[id]; !ok {
		return entityErr.ErrNotFound
	}

	delete(c.Base, id)

	return nil

}

func TestService(t *testing.T) {
	service := New(&clienteRepoMock)

	t.Run("create cliente", func(t *testing.T) {
		c, _ := entities.New(uuid.Nil, "Fulano", "11111111111", "outro@email.com", true)
		cUUID, err := service.Create(*c)
		if err != nil {
			t.Errorf("should not have errors, got: %s", err)
		}
		if _, err := entities.StringToID(cUUID.String()); err != nil {
			t.Errorf("should not return errors, got: %s", err)
		}
	})

	t.Run("creating cliente with existent CPF", func(t *testing.T) {
		c, _ := entities.New(uuid.Nil, "Fulano", existentClientCPF, "outro@email.com", true)
		_, err := service.Create(*c)
		if !errors.Is(err, entityErr.ErrClienteAlreadyExistsForCPF) {
			t.Errorf("want: %s, got: %s", entityErr.ErrClienteAlreadyExistsForCPF, err)
		}
		if err == nil {
			t.Error("should have return error")
		}
	})

	t.Run("creating cliente with existent Email", func(t *testing.T) {
		c, _ := entities.New(uuid.Nil, "Fulano", "22222222222", existentClientEmail, true)
		_, err := service.Create(*c)
		if !errors.Is(err, entityErr.ErrClienteAlreadyExistsForEmail) {
			t.Errorf("want: %s, got: %s", entityErr.ErrClienteAlreadyExistsForEmail, err)
		}
		if err == nil {
			t.Error("should have return error")
		}
	})

	t.Run("creating cliente with invalid CPF", func(t *testing.T) {
		_, err := entities.New(uuid.Nil, "Fulano", "cpfErrado", existentClientEmail, true)
		if !errors.Is(err, entityErr.ErrInvalidCPF) {
			t.Errorf("want: %s, got: %s", entityErr.ErrInvalidCPF, err)
		}
		if err == nil {
			t.Errorf("should have return error")
		}
	})

	t.Run("creating cliente with invalid Email", func(t *testing.T) {
		_, err := entities.New(uuid.Nil, "Fulano", "12312312312", "emailZuado.com", true)
		if !errors.Is(err, entityErr.ErrInvalidEmail) {
			t.Errorf("want: %s, got: %s", entityErr.ErrInvalidEmail, err)
		}
		if err == nil {
			t.Errorf("should have return error")
		}
	})

	t.Run("listing clientes", func(t *testing.T) {
		c, err := service.List()
		if err != nil {
			t.Errorf("should not have any errors, got: %s", err)
		}
		if len(c) != 2 {
			t.Errorf("should have return 2 clientes, got: %d", len(c))
		}
	})

	t.Run("getting inexistent cliente by id", func(t *testing.T) {
		cUUID, _ := entities.StringToID("db6c3a54-541f-472c-8810-13508c930aaa")
		c, err := service.GetClienteById(cUUID)
		if !errors.Is(err, entityErr.ErrNotFound) {
			t.Errorf("should have not found any cliente, got: %s", c.Id())
		}
		if c != nil {
			t.Errorf("should have not found any cliente, got: %s", c.Id())
		}
	})

	t.Run("getting existent cliente by id", func(t *testing.T) {
		cUUID, _ := entities.StringToID(existentClientID)
		c, err := service.GetClienteById(cUUID)
		if err != nil {
			t.Errorf("should have not return any error, got: %s", err)
		}
		if c.Id() != cUUID {
			t.Errorf("should return client with id: %s, got: %s", cUUID, c.Id())
		}
	})

	t.Run("getting inexistent cliente by CPF", func(t *testing.T) {
		c, err := service.GetClienteByCPF("98765432112")
		if !errors.Is(err, entityErr.ErrNotFound) {
			t.Errorf("should have not return any error, got: %s", err)
		}
		if c != nil {
			t.Errorf("should have not return any client, got: %s", c.Id())
		}
	})

	t.Run("getting existent cliente by CPF", func(t *testing.T) {
		c, err := service.GetClienteByCPF(existentClientCPF)
		if err != nil {
			t.Errorf("should have found cliente, got error: %s", err)
		}
		if c.CPF() != existentClientCPF {
			t.Errorf("should have found cliente with the same CPF, want: %s, got: %s", existentClientCPF, c.CPF())
		}
	})

	t.Run("getting inexistent cliente by Email", func(t *testing.T) {
		c, err := service.GetClienteByEmail("hello@email.com")
		if !errors.Is(err, entityErr.ErrNotFound) {
			t.Errorf("should have not return any error, got: %s", err)
		}
		if c != nil {
			t.Errorf("should have not return any client, got: %s", c.Id())
		}
	})

	t.Run("getting existent cliente by Email", func(t *testing.T) {
		c, err := service.GetClienteByEmail(existentClientEmail)
		if err != nil {
			t.Errorf("should have found cliente, got error: %s", err)
		}
		if c.Email() != existentClientEmail {
			t.Errorf("should have found cliente with the same Email, want: %s, got: %s", existentClientEmail, c.Email())
		}
	})

	t.Run("updating non existing cliente", func(t *testing.T) {
		c, _ := entities.New(uuid.Nil, "Fulano", "84738941021", existentClientEmail, true)
		if err := service.Update(*c); err != nil {
			if !errors.Is(err, entityErr.ErrNotFound) {
				t.Error("should not have found cliente")
			}
		} else {
			t.Errorf("should have return error: %s", entityErr.ErrNotFound)
		}
	})

	t.Run("updating existing cliente", func(t *testing.T) {
		cUuid, _ := entities.StringToID(existentClientID)
		c, _ := service.GetClienteById(cUuid)
		c2, _ := entities.New(c.Id(), c.Name(), c.CPF(), "outro2@email.com", c.Active())
		if err := service.Update(*c2); err != nil {
			t.Errorf("should have not return errors, got: %s", err)
		}
	})

	t.Run("deleting non existing cliente", func(t *testing.T) {
		if err := service.Remove(entities.NewID()); err != nil {
			if !errors.Is(err, entityErr.ErrNotFound) {
				t.Error("should not have found cliente")
			}
		} else {
			t.Errorf("should have return error: %s", entityErr.ErrNotFound)
		}
	})

	t.Run("deleting existing cliente", func(t *testing.T) {
		cUuid, _ := entities.StringToID(existentClientID)
		if err := service.Remove(cUuid); err != nil {
			t.Errorf("should have not return errors, got: %s", err)
		}
	})

	t.Run("error creating cliente", func(t *testing.T) {
		errCuui, _ := entities.StringToID(clienteIdError)
		c, _ := entities.New(errCuui, "Fulano", clienteCpfError, clienteEmailError, true)
		_, err := service.Create(*c)
		if err == nil {
			t.Errorf("should have return error")
		}
	})

	t.Run("error listing cliente", func(t *testing.T) {
		flagListClienteError = true

		if _, err := service.List(); err == nil {
			t.Errorf("should have return error")
		}

		flagListClienteError = false
	})

	t.Run("error getting cliente by id", func(t *testing.T) {
		cUuid, _ := entities.StringToID(clienteIdError)
		if _, err := service.GetClienteById(cUuid); err == nil {
			t.Errorf("should have return error")
		}
	})

	t.Run("error getting cliente by cpf", func(t *testing.T) {
		if _, err := service.GetClienteByCPF(clienteCpfError); err == nil {
			t.Errorf("should have return error")
		}
	})

	t.Run("error getting cliente by email", func(t *testing.T) {
		if _, err := service.GetClienteByEmail(clienteEmailError); err == nil {
			t.Errorf("should have return error")
		}
	})

	t.Run("error updating cliente", func(t *testing.T) {
		cUuid, _ := entities.StringToID(clienteIdError)
		c, _ := entities.New(cUuid, "Fulano", clienteCpfError, clienteEmailError, true)
		if err := service.Update(*c); err == nil {
			t.Errorf("should have return error")
		}
	})

	t.Run("error removing cliente", func(t *testing.T) {
		cUuid, _ := entities.StringToID(clienteIdError)
		if err := service.Remove(cUuid); err == nil {
			t.Errorf("should have return error")
		}
	})

}
