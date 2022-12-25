package basic_select

type store interface {
	sell() string
}

type buyer struct {
	store1 store
	store2 store
}

func deliver(store store) <-chan string {
	channel := make(chan string)
	go func() {
		channel <- store.sell()
	}()
	return channel
}

func (b buyer) buy() string {
	order1 := deliver(b.store1)
	order2 := deliver(b.store2)

	select {
	case item := <-order1:
		return item
	case item := <-order2:
		return item
	}
}
