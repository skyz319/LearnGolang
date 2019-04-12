//	城市解析器 解析城市页面下用户列表
package parser

import (
	"LearnGolang/ccmouse_go/crawler/engine"
	"fmt"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>`

//	解析城市下的用户
func ParseCity(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityRe)
	all := re.FindAllSubmatch(contents, -1)

	//	生成ParseResult
	result := engine.ParseResult{}
	for _, m := range all {

		fmt.Printf("Name: %s, Url: %s\n", m[2], m[1])
		//	存相应内容 地外转义为string
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}

	fmt.Println("Matches found: ", len(all))

	return result
}
