package parser

import (
	"LearnGolang/ccmouse_go/crawler_Concurrent/engine"
	"LearnGolang/ccmouse_go/crawler_Concurrent/model"
	"regexp"
	"strings"
)

//	用户ID
var idRe = regexp.MustCompile(`<div class="id".*?">ID：([^<]+)</div>`)

//	个人信息1
var infoRe = regexp.MustCompile(`<div class="m-btn .*?>([^<]+)</div>`)

//	择偶条件
var demandRe = regexp.MustCompile(`<div class="m-btn" .*?>([^<]+)</div>`)

//	相片
var photoRe = regexp.MustCompile(`"photoURL":"([^"]+)"`)

func ParseProfile(contents []byte, name, gender, userURL string) engine.ParseResult {

	profile := model.Profile{
		Name:     name,
		Gender:   gender,
		UserInfo: extractString(contents, infoRe),
		MateInfo: extractString(contents, demandRe),
		Photos:   extractString(contents, photoRe),
	}

	//fmt.Printf("UserInfo >> %s\n", userInfo)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     userURL,
				Type:    "zhenai",
				Id:      extractString(idRe),
				Payload: profile,
			},
		},
	}

	//result := engine.ParseResult{
	//	Items: []engine.Item{
	//		URL: userURL,
	//		Type:"zhenai",
	//		Id: extractString(contents, idRe),
	//		Payload: profile,
	//	},
	//}
	return result
}

// TODO: 需要返回单个sting
func extractString(contents []byte, re *regexp.Regexp) []string {

	match := re.FindAllSubmatch(contents, -1)

	var temp []string

	for _, items := range match {

		str := string(items[1])
		//	照片格式中有u002F，需替换
		temp = append(temp, strings.Replace(str, `\u002F`, `/`, -1))
	}

	return temp

}
