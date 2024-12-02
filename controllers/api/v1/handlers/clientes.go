package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/filipeandrade6/fiap-pedeai-clientes/controllers/api/v1/entities"
	entitiesDomain "github.com/filipeandrade6/fiap-pedeai-clientes/domain/entities"
	"github.com/filipeandrade6/fiap-pedeai-clientes/domain/usecases"

	"github.com/go-chi/chi/v5"
)

func HandleListClientes(clienteUC usecases.ClienteUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("cpf") != "" {
			cliente, err := clienteUC.GetClienteByCPF(r.URL.Query().Get("cpf"))
			if err != nil {
				http.Error(w, "Internal Error", http.StatusInternalServerError)
				return
			}

			ClienteResponse(w, cliente)
			return
		} else if r.URL.Query().Get("email") != "" {
			cliente, err := clienteUC.GetClienteByEmail(r.URL.Query().Get("email"))
			if err != nil {
				http.Error(w, "Internal Error", http.StatusInternalServerError)
				return
			}

			ClienteResponse(w, cliente)
			return
		} else {
			clientes, _ := clienteUC.List()
			var cOut []*entities.Cliente
			for _, c := range clientes {
				out, _ := entities.FromDomain(c)
				cOut = append(cOut, out)
			}

			jEncode := json.NewEncoder(w)
			err := jEncode.Encode(cOut)
			if err != nil {
				http.Error(w, "Internal Error", http.StatusInternalServerError)
				return
			}
		}
	}
}

func HandleGetSingleCliente(clienteUC usecases.ClienteUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		uuid, err := entitiesDomain.StringToID(id)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		c, err := clienteUC.GetClienteById(uuid)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		ClienteResponse(w, c)
		return
	}
}

func HandleCreateCliente(clienteUC usecases.ClienteUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := ClienteDecode(r)
		cDomain, err := c.ToDomain()
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
		uuid, err := clienteUC.Create(*cDomain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(uuid.String()))
	}
}

func HandleUpdateCliente(clienteUC usecases.ClienteUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		uuid, err := entitiesDomain.StringToID(id)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		c, err := ClienteDecode(r)
		c.ID = uuid
		cDomain, err := c.ToDomain()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = clienteUC.Update(*cDomain); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func HandleRemoveCliente(clienteUC usecases.ClienteUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		uuid, err := entitiesDomain.StringToID(id)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		if err := clienteUC.Remove(uuid); err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func ClienteResponse(w http.ResponseWriter, c *entitiesDomain.Cliente) {
	cOut, err := entities.FromDomain(c)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	jEncode := json.NewEncoder(w)
	_ = jEncode.Encode(cOut)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
}

func ClienteDecode(r *http.Request) (*entities.Cliente, error) {
	var c entities.Cliente
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
