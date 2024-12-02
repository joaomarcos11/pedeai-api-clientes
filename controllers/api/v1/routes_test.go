package v1

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/filipeandrade6/fiap-pedeai-clientes/controllers/api/v1/entities"
	domainEntities "github.com/filipeandrade6/fiap-pedeai-clientes/domain/entities"
	entityErr "github.com/filipeandrade6/fiap-pedeai-clientes/domain/errors"

	"github.com/google/uuid"
)

var (
	existentClientID    string = "d1e78e30-2023-4f75-bb3f-41a3b4bacd4d"
	existentClientCPF   string = "12312312312"
	existentClientEmail string = "fulano@email.com"
	existentCliente     *domainEntities.Cliente
)

type ClienteUseCaseMock struct {
	Base map[domainEntities.ID]*domainEntities.Cliente
}

var clienteUCMock ClienteUseCaseMock

func init() {
	uuid, _ := domainEntities.StringToID(existentClientID)
	c, _ := domainEntities.New(
		uuid,
		"Fulano",
		existentClientCPF,
		existentClientEmail,
		true,
	)
	uuid2 := domainEntities.NewID()
	c2, _ := domainEntities.New(
		uuid2,
		"Ciclano",
		"98765432112",
		"ciclano@email.com",
		false,
	)

	mock := ClienteUseCaseMock{}
	mock.Base = make(map[domainEntities.ID]*domainEntities.Cliente)
	mock.Base[uuid] = c
	mock.Base[uuid2] = c2

	clienteUCMock = mock
}

func (c *ClienteUseCaseMock) Create(cliente domainEntities.Cliente) (uuid.UUID, error) {
	return domainEntities.NewID(), nil
}

func (c *ClienteUseCaseMock) List() ([]*domainEntities.Cliente, error) {
	return slices.Collect(maps.Values(c.Base)), nil
}

func (c *ClienteUseCaseMock) GetClienteById(id uuid.UUID) (*domainEntities.Cliente, error) {
	uuid, err := domainEntities.StringToID(id.String())
	if err != nil {
		return nil, errors.New("converting uuid")
	}
	cliente, ok := c.Base[uuid]
	if !ok {
		return nil, entityErr.ErrNotFound
	}
	return cliente, nil
}

func (c *ClienteUseCaseMock) GetClienteByCPF(cpf string) (*domainEntities.Cliente, error) {
	for _, v := range c.Base {
		if v.CPF() == cpf {
			return v, nil
		}
	}

	return nil, entityErr.ErrNotFound
}

func (c *ClienteUseCaseMock) GetClienteByEmail(email string) (*domainEntities.Cliente, error) {
	for _, v := range c.Base {
		if v.Email() == email {
			return v, nil
		}
	}

	return nil, entityErr.ErrNotFound
}

func (c *ClienteUseCaseMock) Update(cliente domainEntities.Cliente) error {
	if _, ok := c.Base[cliente.Id()]; !ok {
		return entityErr.ErrNotFound
	}
	c.Base[cliente.Id()] = &cliente
	return nil
}

func (c *ClienteUseCaseMock) Remove(id uuid.UUID) error {
	if _, ok := c.Base[id]; !ok {
		return entityErr.ErrNotFound
	}
	delete(c.Base, id)
	return nil
}

func TestHandlers(t *testing.T) {
	routes := AddRoutes(&clienteUCMock)

	t.Run("list clientes", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/clientes", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		routes.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		clientes := []*entities.Cliente{}
		if err := json.Unmarshal(rr.Body.Bytes(), &clientes); err != nil {
			t.Errorf("unmarshalling json: %s", err)
		}

		if len(clientes) != 2 {
			t.Errorf("should have return list with 2 itens, got: %d", len(clientes))
		}
	})

	t.Run("get cliente by id", func(t *testing.T) {
		req, err := http.NewRequest("GET", fmt.Sprintf("/clientes/%s", existentClientID), nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		routes.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		cliente := entities.Cliente{}
		if err := json.Unmarshal(rr.Body.Bytes(), &cliente); err != nil {
			t.Errorf("unmarshalling json: %s", err)
		}

		if cliente.ID.String() != existentClientID {
			t.Errorf("should have return existent cliente with id: %s, got: %s", existentClientID, cliente.ID)
		}
	})

	t.Run("get cliente by CPF", func(t *testing.T) {
		req, err := http.NewRequest("GET", fmt.Sprintf("/clientes?cpf=%s", existentClientCPF), nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		routes.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		cliente := entities.Cliente{}
		if err := json.Unmarshal(rr.Body.Bytes(), &cliente); err != nil {
			t.Errorf("unmarshalling json: %s", err)
		}

		if cliente.CPF != existentClientCPF {
			t.Errorf("should have return existent cliente with cpf: %s, got: %s", existentClientCPF, cliente.CPF)
		}
	})

	t.Run("get cliente by Email", func(t *testing.T) {
		req, err := http.NewRequest("GET", fmt.Sprintf("/clientes?email=%s", existentClientEmail), nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		routes.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		cliente := entities.Cliente{}
		if err := json.Unmarshal(rr.Body.Bytes(), &cliente); err != nil {
			t.Errorf("unmarshalling json: %s", err)
		}

		if cliente.CPF != existentClientCPF {
			t.Errorf("should have return existent cliente with email: %s, got: %s", existentClientEmail, cliente.Email)
		}
	})

	t.Run("create cliente", func(t *testing.T) {
		cliente := entities.Cliente{
			Name:   "Fulano",
			CPF:    "83483483423",
			Email:  "outro@email.com",
			Active: false,
		}

		c, err := json.Marshal(cliente)
		if err != nil {
			t.Fatalf("marshalling cliente into json, error: %s", err)
		}
		b := bytes.NewBuffer(c)

		req, err := http.NewRequest("POST", "/clientes", b)
		if err != nil {
			t.Fatalf("creating request, error: %s", err)
		}

		rr := httptest.NewRecorder()
		routes.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}
	})

	t.Run("update cliente", func(t *testing.T) {
		cUuid, _ := domainEntities.StringToID(existentClientID)
		clienteMock := clienteUCMock.Base[cUuid]
		cliente := entities.Cliente{
			Name:   "Ciclano2",
			CPF:    "74374374326",
			Email:  "outro2@email.com",
			Active: false,
		}

		c, err := json.Marshal(cliente)
		if err != nil {
			t.Fatalf("marshalling cliente into json, error: %s", err)
		}
		b := bytes.NewBuffer(c)

		req, err := http.NewRequest("PUT", fmt.Sprintf("/clientes/%s", clienteMock.Id()), b)
		if err != nil {
			t.Fatalf("creating request, error: %s", err)
		}

		rr := httptest.NewRecorder()
		routes.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
		}

		clienteMockUpdated := clienteUCMock.Base[cUuid]
		if clienteMockUpdated.CPF() != cliente.CPF {
			t.Errorf("should have updated cliente")
		}
	})

	t.Run("remove cliente", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/clientes/%s", existentClientID), nil)
		if err != nil {
			t.Fatalf("creating request, error: %s", err)
		}

		rr := httptest.NewRecorder()
		routes.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
		}

		cUuid, _ := domainEntities.StringToID(existentClientID)
		if _, ok := clienteUCMock.Base[cUuid]; ok {
			t.Errorf("should have deleted cliente")
		}
	})
}
