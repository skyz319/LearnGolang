// 使用goroutine
package scheduler

import "LearnGolang/ccmouse_go/crawler_Concurrent/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {

	//	为每个Request开一个goroutine, 以便往worker分发
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) ConfigureMasterWorkChan(c chan engine.Request) {

	s.workerChan = c
}
