package executor

import (
	"sync"
)

type Service struct {
	jobChan chan func()
}

func (s *Service) Start(workersCnt int, setup func()) {
	jobChan := make(chan func())
	s.jobChan = jobChan
	wg := sync.WaitGroup{}
	wg.Add(workersCnt)

	for i := 0; i < workersCnt; i++ {
		go worker(setup, &wg, jobChan)
	}

	wg.Wait()
}

func (s Service) Run(job func()) {
	s.jobChan <- job
}

func worker(setup func(), wg *sync.WaitGroup, jobChan chan func()) {
	setup()
	wg.Done()
	for {
		job := <-jobChan
		job()
	}
}
