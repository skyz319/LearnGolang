//	驱动引擎，维护队列
package engine

import (
	"LearnGolang/ccmouse_go/crawler/fetcher"
	"log"
)

func Ren(seeds ...Request) {

	//	请求队列
	var requestsQueue []Request

	for _, r := range seeds {

		requestsQueue = append(requestsQueue, r)
	}

	for len(requestsQueue) > 0 {

		//	取出第一个request并抓取其内容
		r := requestsQueue[0]
		requestsQueue = requestsQueue[1:]

		log.Printf("engin.go >> Fetching %s", r.Url)

		//	抓取内容
		body, err := fetcher.Fetch(r.Url)
		if err != nil {

			//	忽略取出的错误Url
			log.Printf("engin.go >> Fetcher: error "+"fetching url %s: %v\n", r.Url, err)
			continue
		}

		//	将任务中的request添加进队列
		parseResult := r.ParserFunc(body)
		//	...将requests中所有内容展开并添加到指定队列
		requestsQueue = append(requestsQueue, parseResult.Requests...)

		for _, item := range parseResult.Items {

			log.Printf("engin.go >> Got item %v ", item)
		}
	}
}
