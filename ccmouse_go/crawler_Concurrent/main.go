package main

import (
	"LearnGolang/ccmouse_go/crawler_Concurrent/engine"
	"LearnGolang/ccmouse_go/crawler_Concurrent/persist"
	"LearnGolang/ccmouse_go/crawler_Concurrent/scheduler"
	"LearnGolang/ccmouse_go/crawler_Concurrent/zhenai/parser"
)

const BaseURL = "http://www.zhenai.com/zhenghun"

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 50,
		ItemChan:    persist.ItemSaver(),
	}

	//e.Ren(engine.Request{
	//	Url:        BaseURL,
	//	ParserFunc: parser.ParseCityList,
	//})

	//	爬取单一城市
	e.Ren(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/chengdu",
		ParserFunc: parser.ParseCity,
	})
}
