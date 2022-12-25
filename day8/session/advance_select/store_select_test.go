package advance_select

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockStore struct {
	name      string
	available bool
}

func (m mockStore) sell(item string) string {
	return "order for " + item
}

func (m mockStore) isAvailable(item string) bool {
	return m.available
}

func TestDeliver(t *testing.T) {
	mockStore := mockStore{name: "store", available: false}

	_, err := buyer{}.buy(mockStore, "apple")

	assert.Error(t, err)
	assert.Equal(t, "apple not available", err.Error())
}
