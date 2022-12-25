package executor

import (
	"errors"
	"sync"
)

type Service struct {
	jobChan  chan func()
	isClosed bool
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

func (s Service) Run(job func()) error {
	if s.isClosed {
		return errors.New("Closed")
	}
	s.jobChan <- job
	return nil
}

func (s *Service) Close() {
	close(s.jobChan)
	s.isClosed = true
}

func (s *Service) RunBatch(jobs []func()) error {
	if s.isClosed {
		return errors.New("Closed")
	}
	for i := 0; i < len(jobs); i++ {
		err := s.Run(jobs[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func worker(setup func(), wg *sync.WaitGroup, jobChan chan func()) {
	setup()
	wg.Done()
	for {
		job, ok := <-jobChan
		// ok is true if it is closed channel
		if !ok {
			break
		}
		job()
	}
}
