package parser

import (
	"LearnGolang/ccmouse_go/crawler_Concurrent/engine"
	"regexp"
	"strings"
)

type UserInfo struct {
	nick   string   //	昵称
	gender string   //	性别
	info   []string //	用户信息
	demand []string //	择偶条件
	photo  []string //	照片列表
}

//	个人信息1
var infoRe = regexp.MustCompile(`<div class="m-btn .*?>([^<]+)</div>`)

//	择偶条件
var demandRe = regexp.MustCompile(`<div class="m-btn" .*?>([^<]+)</div>`)

//	相片
var photoRe = regexp.MustCompile(`"photoURL":"([^"]+)"`)

func ParseProfile(contents []byte, name, gender string) engine.ParseResult {

	var userInfo UserInfo
	userInfo.nick = name
	userInfo.gender = gender
	userInfo.info = extractString(contents, infoRe)
	userInfo.demand = extractString(contents, demandRe)
	userInfo.photo = extractString(contents, photoRe)

	//fmt.Printf("UserInfo >> %s\n", userInfo)

	result := engine.ParseResult{
		Items: []interface{}{userInfo},
	}
	return result
}

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
