package executor

import "sync"

type Service struct {
}

func (s Service) Start(workersCnt int, setup func()) {
	wg := sync.WaitGroup{}
	wg.Add(workersCnt)
	for i := 0; i < workersCnt; i++ {
		go worker(setup, &wg)
	}

	wg.Wait()
}

func worker(setup func(), wg *sync.WaitGroup) {
	setup()
	wg.Done()
}
