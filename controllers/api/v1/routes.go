package v1

import (
	"github.com/filipeandrade6/fiap-pedeai-clientes/controllers/api/v1/handlers"
	"github.com/filipeandrade6/fiap-pedeai-clientes/domain/usecases"

	"github.com/go-chi/chi/v5"
)

func AddRoutes(clienteUC usecases.ClienteUseCase) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/clientes", func(r chi.Router) {
		r.Get("/", handlers.HandleListClientes(clienteUC))
		r.Get("/{id}", handlers.HandleGetSingleCliente(clienteUC))
		r.Post("/", handlers.HandleCreateCliente(clienteUC))
		r.Put("/{id}", handlers.HandleUpdateCliente(clienteUC))
		r.Delete("/{id}", handlers.HandleRemoveCliente(clienteUC))
	})

	return r
}
