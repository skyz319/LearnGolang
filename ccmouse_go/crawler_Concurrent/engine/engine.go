//	驱动引擎，维护队列
package engine

import (
	"LearnGolang/ccmouse_go/crawler_Concurrent/fetcher"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler //	调度器
	WorkerCount int       //	worker 数量
}

//	调度者接口
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkChan(chan Request)
}

func (e *ConcurrentEngine) Ren(seeds ...Request) {

	//	输入/输出 队列
	in := make(chan Request)
	out := make(chan ParseResult)
	//	worker队列
	e.Scheduler.ConfigureMasterWorkChan(in)

	for i := 0; i < e.WorkerCount; i++ {

		//	创建worker
		e.createWorker(in, out)
	}

	for _, r := range seeds {
		//	将任务放入调度器
		e.Scheduler.Submit(r)
	}

	//	收取返回的数据
	for {

		result := <-out

		//	计数
		itemCount := 0
		for _, item := range result.Items {

			log.Printf("Got item #%d: %s\n", itemCount, item)
		}

		//	将返回数据中的request 送入调度者
		for _, req := range result.Requests {

			e.Scheduler.Submit(req)
		}
	}
}

func (e ConcurrentEngine) createWorker(in chan Request, out chan ParseResult) {

	go func() {

		for {

			request := <-in

			result, err := e.worker(request)

			if err != nil {

				continue
			}

			out <- result
		}
	}()
}

func (e ConcurrentEngine) worker(r Request) (ParseResult, error) {

	log.Printf("engin.go >> Fetching %s", r.Url)

	//	抓取内容
	body, err := fetcher.Fetch(r.Url)
	if err != nil {

		//	忽略取出的错误Url
		log.Printf("engin.go >> Fetcher: error "+"fetching url %s: %v\n", r.Url, err)

		return ParseResult{}, err
	}

	//	将任务中的request添加进队列
	return r.ParserFunc(body), nil

}
