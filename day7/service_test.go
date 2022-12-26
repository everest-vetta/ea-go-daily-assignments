package executor

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	s := Service{}

	expectedWorkerCnt := 3
	var actualCnt int32 = 0
	setUp := func() {
		atomic.AddInt32(&actualCnt, 1)
	}

	s.Start(expectedWorkerCnt, setUp)
	assert.Equal(t, expectedWorkerCnt, int(actualCnt))
}

func TestRun(t *testing.T) {
	s := Service{}

	expectedWorkerCnt := 3
	expectedJobCnt := 4
	var actualCnt int32 = 0
	s.Start(expectedWorkerCnt, func() {})

	wg := sync.WaitGroup{}
	wg.Add(expectedJobCnt)
	job := func() {
		wg.Done()
		atomic.AddInt32(&actualCnt, 1)
	}
	for i := 0; i < expectedJobCnt; i++ {
		s.Run(job)
	}

	wg.Wait()
	assert.Equal(t, int(actualCnt), expectedJobCnt)
}

func TestClose(t *testing.T) {
	s := Service{}
	expectedWorkerCnt := 3
	s.Start(expectedWorkerCnt, func() {})

	s.Close()

	isExecuted := false
	job := func() {
		isExecuted = true

	}
	err := s.Run(job)
	assert.Error(t, err)
	assert.False(t, isExecuted)
}

func TestRunBatch(t *testing.T) {
	s := Service{}

	expectedWorkerCnt := 3
	s.Start(expectedWorkerCnt, func() {})

	var actualCnt int32 = 0
	expectedJobArraySize := 5

	wg := sync.WaitGroup{}
	wg.Add(expectedJobArraySize)
	job := func() {
		atomic.AddInt32(&actualCnt, 1)
		wg.Done()
	}

	var jobs []func()
	for i := 0; i < expectedJobArraySize; i++ {
		jobs = append(jobs, job)
	}

	s.RunBatch(jobs)
	wg.Wait()
	assert.Equal(t, expectedJobArraySize, int(actualCnt))
}

func TestWorker(t *testing.T) {
	workerCnt := 2
	ch := make(chan func(), 5)
	expectedJobCnt := 5
	var actualCnt int32 = 0

	wg2 := sync.WaitGroup{}
	wg2.Add(workerCnt)

	wg1 := sync.WaitGroup{}
	wg1.Add(expectedJobCnt)

	job := func() {
		wg1.Done()
		atomic.AddInt32(&actualCnt, 1)
	}

	for i := 0; i < workerCnt; i++ {
		go Worker(func() {}, &wg2, ch)
	}

	for i := 0; i < expectedJobCnt; i++ {
		ch <- job
	}

	wg1.Wait()
	close(ch)
	wg2.Wait()

	assert.Equal(t, expectedJobCnt, int(actualCnt))
}

// func TestWritingToClosedChannel(t *testing.T) {
// 	ch := make(chan func())
// 	close(ch)
// 	job := func() {

// 	}

// 	ch <- job
// }

func TestListeningToClosedChannel(t *testing.T) {
	ch := make(chan func())
	close(ch)
	<-ch
}
