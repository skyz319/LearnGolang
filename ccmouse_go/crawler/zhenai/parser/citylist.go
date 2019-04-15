//	城市列表解析器
package parser

import (
	"LearnGolang/ccmouse_go/crawler/engine"
	"fmt"
	"regexp"
)

//	[^>]* 非右括号的1个或多个字符
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]+>([^<]+)</a>`

//	获取城市列表
func ParseCityList(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityListRe)
	all := re.FindAllSubmatch(contents, -1)

	//	生成ParseResult
	result := engine.ParseResult{}
	//	限制城市数量
	limit := engine.CityNum
	for _, m := range all {
		fmt.Printf("citylist.go >> City: %s, URL: %s\n", m[2], m[1])

		//	存相应内容 地外转义为string
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})

		//	限制抓取的城市数量
		if engine.Limit {
			limit--
			if limit == 0 {
				break
			}
		}
	}

	fmt.Println("citylist.go >> Matches found: ", len(all))

	return result
}
