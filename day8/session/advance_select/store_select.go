package advance_select

import (
	"fmt"
	"time"
)

type store interface {
	sell(item string) string
	isAvailable(items string) bool
}

type buyer struct {
}

func deliver(store store, item string) <-chan string {
	channel := make(chan string)
	go func() {
		if store.isAvailable(item) {
			channel <- store.sell(item)
		}
	}()
	return channel
}

func (b buyer) buy(store store, itemName string) (string, error) {
	order := deliver(store, itemName)
	select {
	case item := <-order:
		return item, nil
	case <-time.After(10 * time.Millisecond):
		return "", fmt.Errorf("%s not available", itemName)
	}
}
