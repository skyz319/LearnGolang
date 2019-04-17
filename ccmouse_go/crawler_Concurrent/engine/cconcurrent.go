//	驱动引擎，维护队列
package engine

import (
	"LearnGolang/ccmouse_go/crawler_Concurrent/fetcher"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler //	调度器
	WorkerCount int       //	worker 数量
	ItemChan    chan interface{}
}

//	调度者接口
type Scheduler interface {
	ReadNotifier
	Submit(Request)
	WorkerChan() chan Request //	由scheduler分配chan
	Run()
}

type ReadNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Ren(seeds ...Request) {

	//	输出 队列
	out := make(chan ParseResult)
	//	worker队列
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {

		//	创建worker
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		//	将任务放入调度器
		e.Scheduler.Submit(r)
	}

	//	收取返回的数据
	for {

		result := <-out
		for _, item := range result.Items {

			go func() {
				//	将数据放入ItemSaver中
				e.ItemChan <- item
			}()
		}

		//	将返回数据中的request 送入调度者
		for _, req := range result.Requests {

			e.Scheduler.Submit(req)
		}
	}
}

func (e ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadNotifier) {

	go func() {

		for {

			ready.WorkerReady(in)

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

	//log.Printf("concurrent.go >> Fetching %s", r.Url)

	//	抓取内容
	body, err := fetcher.Fetch(r.Url)
	if err != nil {

		//	忽略取出的错误Url
		log.Printf("concurrent.go >> Fetcher: error "+"fetching url %s: %v\n", r.Url, err)

		return ParseResult{}, err
	}

	//	将任务中的request添加进队列
	return r.ParserFunc(body), nil

}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {

	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true

	return false
}
