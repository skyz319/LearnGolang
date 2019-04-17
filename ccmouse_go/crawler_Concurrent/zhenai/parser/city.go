//	城市解析器 解析城市页面下用户列表
package parser

import (
	"LearnGolang/ccmouse_go/crawler_Concurrent/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<div class="content">.*?<a href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>.*?<td width="180"><span class="grayL">性别：</span>([^<]+)</td>.*?</div>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

//	解析城市下的用户
func ParseCity(contents []byte) engine.ParseResult {

	all := profileRe.FindAllSubmatch(contents, -1)

	//	生成ParseResult
	result := engine.ParseResult{}
	//	限制抓取的用户数量
	limit := engine.UserNum

	for _, m := range all {
		name := string(m[2])
		gender := string(m[3])
		userURL := string(m[1])
		//	排除指定性别
		if gender == engine.Exclusion {
			break
		}

		//fmt.Printf("city.go >> Name:Name %s, Url: %s\n", m[2], m[1])
		//	存相应内容 地外转义为string
		//result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url: userURL,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name, gender, userURL)
			},
		})

		if engine.LimitUser {
			limit--
			if limit == 0 {
				break
			}
		}

	}

	//fmt.Println("city.go >> Matches found: ", len(all))

	//	添加其它城市
	matches := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {

		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
