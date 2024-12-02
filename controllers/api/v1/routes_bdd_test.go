package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/filipeandrade6/fiap-pedeai-clientes/controllers/api/v1/entities"
	domainEntities "github.com/filipeandrade6/fiap-pedeai-clientes/domain/entities"
)

// BDD
// Feature: Get cliente searching by ID
// Scenario: Successfully retrieve cliente information searching by ID
func TestBDD(t *testing.T) {
	routes := AddRoutes(&clienteUCMock)

	t.Run("get cliente by id", func(t *testing.T) {

		// Given the cliente with ID "d1e78e30-2023-4f75-bb3f-41a3b4bacd4d" exists in the system
		uuid, _ := domainEntities.StringToID("d1e78e30-2023-4f75-bb3f-41a3b4bacd4d")

		// When the GET request to the endpoint "/{id}" with the ID "d1e78e30-2023-4f75-bb3f-41a3b4bacd4d" is sended
		req, err := http.NewRequest("GET", fmt.Sprintf("/clientes/%s", uuid), nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		routes.ServeHTTP(rr, req)

		// Then the http status code response should be 200 "Status OK"
		if status := rr.Code; status != 200 {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		// And the http response body should contain the cliente details
		cliente := entities.Cliente{}
		if err := json.Unmarshal(rr.Body.Bytes(), &cliente); err != nil {
			t.Errorf("unmarshalling json: %s", err)
		}
		if cliente.ID.String() != existentClientID {
			t.Errorf("should have return existent cliente with id: %s, got: %s", existentClientID, cliente.ID)
		}
	})
}
