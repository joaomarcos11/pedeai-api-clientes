package postgresql

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/filipeandrade6/fiap-pedeai-clientes/domain/entities"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestWithPostgres(t *testing.T) {
	ctx := context.Background()

	dbName := "pedeai"
	dbUser := "pedeai"
	dbPassword := "senha1ABC"

	postgresContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithInitScripts("./../../../deployments/init.sql"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	defer func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}

	repo, err := New(context.Background(), Config{
		Host:       "localhost",
		Port:       "5432",
		User:       dbUser,
		Password:   dbPassword,
		Name:       dbName,
		DisableTLS: true,
	})
	if err != nil {
		t.Errorf("should not return any error, got: %s", err)
	}

	usedUuid := entities.NewID()
	c, _ := entities.New(usedUuid, "Fulano", "12312312312", "fulanoZZZ@email.com", true)

	if err := repo.db.DeleteAllCliente(context.Background()); err != nil {
		t.Errorf("cleaning db, got error: %s", err)
	}

	t.Run("listing empty cliente", func(t *testing.T) {
		cs, err := repo.List()
		if len(cs) != 0 {
			t.Error("should return empty list")
		}
		if err != nil {
			t.Errorf("should not return any error, got: %s", err)
		}
	})

	t.Run("create cliente", func(t *testing.T) {
		err = repo.Create(*c)
		if err != nil {
			t.Errorf("should not have return any error, got: %s", err)
		}

	})

	t.Run("list cliente", func(t *testing.T) {
		_, err := repo.List()
		if err != nil {
			t.Errorf("should not have return any error, got: %s", err)
		}

	})

	t.Run("get cliente by id", func(t *testing.T) {
		_, err = repo.GetClienteById(usedUuid)
		if err != nil {
			t.Errorf("should not have return any error, got: %s", err)
		}

	})

	t.Run("get cliente by cpf", func(t *testing.T) {
		_, err := repo.GetClienteByCPF("12312312312")
		if err != nil {
			t.Errorf("should not have return any error, got: %s", err)
		}

	})

	t.Run("get cliente by email", func(t *testing.T) {
		_, err = repo.GetClienteByEmail("fulanoZZZ@email.com")
		if err != nil {
			t.Errorf("should not have return any error, got: %s", err)
		}

	})

	t.Run("update cliente", func(t *testing.T) {
		c2, _ := entities.New(c.Id(), "Ciclano", c.CPF(), c.Email(), false)
		err = repo.Update(*c2)
		if err != nil {
			t.Errorf("should not have return any error, got: %s", err)
		}

	})

	t.Run("remove cliente", func(t *testing.T) {
		err = repo.Remove(c.Id())
		if err != nil {
			t.Errorf("should not have return any error, got: %s", err)
		}

	})

}
