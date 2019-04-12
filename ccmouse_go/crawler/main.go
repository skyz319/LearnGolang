package main

import (
	"LearnGolang/ccmouse_go/crawler/engine"
	"LearnGolang/ccmouse_go/crawler/zhenai/parser"
)

const BaseURL = "http://www.zhenai.com/zhenghun"

func main() {

	engine.Ren(engine.Request{
		Url:        BaseURL,
		ParserFunc: parser.ParseCityList,
	})
}
