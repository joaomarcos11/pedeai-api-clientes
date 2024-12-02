package entities

import (
	"testing"

	"github.com/google/uuid"
)

func TestEntity(t *testing.T) {

	t.Run("creating uuid", func(t *testing.T) {
		NewUUID := NewID()
		if err := uuid.Validate(NewUUID.String()); err != nil {
			t.Error("invalid uuid")
		}
	})

	t.Run("validating incorrect uuid", func(t *testing.T) {
		invalidUUID := "6bb9cfdc-32d9-49a3-99d9-cdbf910f2ecz"
		if _, err := StringToID(invalidUUID); err == nil {
			t.Error("should have return error for invlaid uuid parsing")
		}
	})

	t.Run("validating correct uuid", func(t *testing.T) {
		validUUID := "6bb9cfdc-32d9-49a3-99d9-cdbf910f2ec7"
		if _, err := StringToID(validUUID); err != nil {
			t.Error("should have return error for invlaid uuid parsing")
		}
	})

}
