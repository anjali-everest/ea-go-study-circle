package executor

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"sync/atomic"
	"testing"
)

func TestStart(t *testing.T) {
	s := Service{}
	expCnt := 3
	var actualCnt int32 = 0

	setup := func() {
		atomic.AddInt32(&actualCnt, 1)
	}

	s.Start(3, setup)

	assert.Equal(t, expCnt, int(actualCnt))
}

func TestRun(t *testing.T) {
	s := Service{}
	s.Start(3, func() {})
	expCnt := 4
	var actualCnt int32 = 0

	wg := sync.WaitGroup{}
	wg.Add(expCnt)
	job := func() {
		wg.Done()
		atomic.AddInt32(&actualCnt, 1)
	}
	for i := 0; i < expCnt; i++ {
		s.Run(job)
	}
	wg.Wait()
	assert.Equal(t, expCnt, int(actualCnt))
}
