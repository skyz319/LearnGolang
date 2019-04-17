//	队列版
package scheduler

import "LearnGolang/ccmouse_go/crawler_Concurrent/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	//	Worker队列
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {

	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {

	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {

	s.workerChan <- w
}

//	总控
func (s *QueuedScheduler) Run() {

	//	生成chan
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)

	go func() {

		//	请求队列及worker队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {

			var activeRequest engine.Request
			var activeWoker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {

				activeRequest = requestQ[0]
				activeWoker = workerQ[0]
			}

			//	队列添加，使用Switch是为方便先有的数据
			select {
			case r := <-s.requestChan: //	有任务进来时，将任务加入任务队列
				requestQ = append(requestQ, r)
			case w := <-s.workerChan: //	有空闲worker时，加入worker队列
				workerQ = append(workerQ, w)
			case activeWoker <- activeRequest: //	将请求送给空闲worker
				//	移除队列中已添加的请求和worker
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]

			}

		}
	}()
}
