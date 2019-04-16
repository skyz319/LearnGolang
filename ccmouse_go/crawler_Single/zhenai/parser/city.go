//	城市解析器 解析城市页面下用户列表
package parser

import (
	"LearnGolang/ccmouse_go/crawler_Concurrent/engine"
	"fmt"
	"regexp"
)

//const cityRe = `<a href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>`
var cityRe = regexp.MustCompile(`<div class="content">.*?<a href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>.*?<td width="180"><span class="grayL">性别：</span>([^<]+)</td>.*?</div>`)

//	解析城市下的用户
func ParseCity(contents []byte) engine.ParseResult {

	all := cityRe.FindAllSubmatch(contents, -1)

	//	生成ParseResult
	result := engine.ParseResult{}
	//	限制抓取的用户数量
	limit := engine.UserNum

	for _, m := range all {
		name := string(m[2])
		gender := string(m[3])
		//	排除指定性别
		if gender == engine.Exclusion {
			break
		}

		fmt.Printf("city.go >> Name: %s, Url: %s\n", m[2], m[1])
		//	存相应内容 地外转义为string
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name, gender)
			},
		})

		if engine.Limit {
			limit--
			if limit == 0 {
				break
			}
		}

	}

	fmt.Println("city.go >> Matches found: ", len(all))

	return result
}
