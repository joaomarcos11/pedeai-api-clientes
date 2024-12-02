package entities

import (
	"errors"
	"testing"

	entityErr "github.com/filipeandrade6/fiap-pedeai-clientes/domain/errors"
)

func TestCliente(t *testing.T) {

	testId := NewID()
	testName := "Fulano"
	testCpf := "12312312312"
	testEmail := "fulano@email.com"
	testActive := false

	c, _ := New(testId, testName, testCpf, testEmail, testActive)

	t.Run("creating cliente", func(t *testing.T) {
		assertCorrectId(t, c.Id(), testId)
		assertCorrectString(t, c.Name(), testName)
		assertCorrectString(t, c.CPF(), testCpf)
		assertCorrectString(t, c.Email(), testEmail)
		assertCorrectBoolean(t, c.Active(), testActive)
	})

	t.Run("empty name", func(t *testing.T) {
		_, err := New(testId, "", testCpf, testEmail, testActive)
		if !errors.Is(err, entityErr.ErrNameRequired) {
			t.Errorf("wanted %s error got %s", entityErr.ErrNameRequired, err)
		}
		if err == nil {
			t.Errorf("wanted %s error got nil", entityErr.ErrNameRequired)
		}
	})

	t.Run("name too short", func(t *testing.T) {
		_, err := New(testId, "ab", testCpf, testEmail, testActive)
		if !errors.Is(err, entityErr.ErrNameTooShort) {
			t.Errorf("wanted %s error got %s", entityErr.ErrNameTooShort, err)
		}
		if err == nil {
			t.Errorf("wanted %s error got nil", entityErr.ErrNameTooShort)
		}
	})

	t.Run("wrong cpf format", func(t *testing.T) {
		test2Cpf := "123123123123"

		_, err := New(testId, testName, test2Cpf, testEmail, testActive)
		if !errors.Is(err, entityErr.ErrInvalidCPF) {
			t.Errorf("wanted %s error got %s", entityErr.ErrInvalidCPF, err)
		}
		if err == nil {
			t.Errorf("wanted %s error got nil", entityErr.ErrInvalidCPF)
		}
	})

	t.Run("wrong email format", func(t *testing.T) {
		test2Email := "fulano.c"

		_, err := New(testId, testName, testCpf, test2Email, testActive)
		if !errors.Is(err, entityErr.ErrInvalidEmail) {
			t.Errorf("wanted %s error got %#v", entityErr.ErrInvalidEmail, err)
		}
		if err == nil {
			t.Errorf("wanted %s error got nil", entityErr.ErrInvalidEmail)
		}
	})
}

func assertCorrectId(t testing.TB, got, want ID) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertCorrectString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertCorrectBoolean(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
