package parser

import (
	"LearnGolang/ccmouse_go/crawler/engine"
	"regexp"
)

//	昵称
var nickRe = regexp.MustCompile(`<h1 class="nickName" .*?>([^<]+)</h1>`)

//	个人信息1
var infoRe1 = regexp.MustCompile(`<div class="m-btn .*?>([^<]+)</div>`)

//	个人信息2
var infoR22 = regexp.MustCompile(`<div class="m-btn pink" .*?>([^<]+)</div>`)

//	择偶条件
var demandRe = regexp.MustCompile(`<div class="m-btn" .*?>([^<]+)</div>`)

//	相片
var photoRe = regexp.MustCompile(`<div .*? href="([^"]+)" class="photoItem z-cursor-big.*?">`)

//	<div class="purple-btns" data-v-bff6f798=""><div class="m-btn purple" data-v-bff6f798="">未婚</div><div class="m-btn purple" data-v-bff6f798="">43岁</div><div class="m-btn purple" data-v-bff6f798="">魔羯座(12.22-01.19)</div><div class="m-btn purple" data-v-bff6f798="">166cm</div><div class="m-btn purple" data-v-bff6f798="">56kg</div><div class="m-btn purple" data-v-bff6f798="">工作地:鞍山铁东区</div><div class="m-btn purple" data-v-bff6f798="">月收入:3千以下</div><div class="m-btn purple" data-v-bff6f798="">美容师</div><div class="m-btn purple" data-v-bff6f798="">中专</div></div>

func ParseProfile(contents []byte) engine.ParseResult {

	//re := regexp.MustCompile(ageRe)
	//match := re.FindAllSubmatch(contents, -1)
	//
	//if match != nil {
	//}

	return engine.ParseResult{}
}

func extractString(contents []byte, re *regexp.Regexp) string {

	match := re.FindSubmatch(contents)

	if len(match) >= 2 {

		return string(match[1])
	} else {

		return ""
	}
}
