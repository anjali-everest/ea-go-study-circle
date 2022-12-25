package executor

import (
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
)

func TestStart(t *testing.T) {
	s := Service{}
	var expCnt int32 = 3
	var actualCnt int32 = 0
	setup := func() {
		atomic.AddInt32(&actualCnt, 1)
	}
	s.Start(3, setup)

	assert.Equal(t, expCnt, actualCnt)
}
