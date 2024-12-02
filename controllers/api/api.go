package api

import (
	"log/slog"
	"net/http"

	v1 "github.com/filipeandrade6/fiap-pedeai-clientes/controllers/api/v1"
	"github.com/filipeandrade6/fiap-pedeai-clientes/domain/usecases"

	"github.com/go-chi/chi/v5"
)

func NewServer(
	logger *slog.Logger,
	clienteUC usecases.ClienteUseCase,
) http.Handler {
	r := chi.NewRouter()

	r.Mount("/v1", v1.AddRoutes(clienteUC))

	return r
}
