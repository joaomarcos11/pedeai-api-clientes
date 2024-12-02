package api

import "testing"

func TestAPI(t *testing.T) {
	t.Run("test API", func(t *testing.T) {
		_ = NewServer(nil, nil)
	})
}
