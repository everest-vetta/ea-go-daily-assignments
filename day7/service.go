package executor

import (
	"errors"
	"sync"
)

type Service struct {
	ch       chan func()
	isClosed bool
}

func (s *Service) Start(workersCnt int, setUp func()) {
	ch := make(chan func(), 5)
	s.ch = ch
	wg := sync.WaitGroup{}
	wg.Add(workersCnt)

	for i := 0; i < workersCnt; i++ {
		go Worker(setUp, &wg, s.ch)
	}
	wg.Wait()
}

func (s Service) Run(job func()) error {
	if s.isClosed {
		return errors.New("channel closed")
	}
	s.ch <- job
	return nil
}

func (s *Service) Close() {
	close(s.ch)
	s.isClosed = true
}

func (s *Service) RunBatch(jobs []func()) {
	for i := range jobs {
		//s.Run(jobs[i])
		s.ch <- jobs[i]
	}
}

func Worker(setUp func(), wg *sync.WaitGroup, ch chan func()) {
	setUp()
	wg.Done()

	for {
		job, ok := <-ch
		if !ok {
			break
		}
		job()
	}
}
