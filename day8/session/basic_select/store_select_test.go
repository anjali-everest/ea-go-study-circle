package basic_select

import (
	"testing"
	"time"
)

type mockStore struct {
	name string
	time time.Duration
}

func (m mockStore) sell() string {
	time.Sleep(m.time)
	return m.name
}

func TestDeliver(t *testing.T) {
	mockStore1 := mockStore{name: "store1", time: 10 * time.Millisecond}
	mockStore2 := mockStore{name: "store2", time: 20 * time.Millisecond}

	buyer := buyer{mockStore1, mockStore2}
	order := buyer.buy()

	if order != mockStore1.name {
		t.Errorf("%s is not an expected store", order)
	}
}
