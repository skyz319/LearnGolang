package main

import (
	"fmt"
	"regexp"
	"strings"
)

const text = `My email is skyz319@gmail.com
email1 is cain.code@qq.com
email2 is 1691164@qq.com
sky.z@tom.com
skyz_319@tom.com
skyz_319@tom.com.cn
`

const text2 = `
<div data-v-5b109fc3="" class="photoBox" style="width: 430px; left: 0px;"><div data-v-5b109fc3="" href="https://photo.zastatic.com/images/photo/440314/1761252184/26087356455159750.jpg" class="photoItem z-cursor-big active"><img data-v-5b109fc3="" src="https://photo.zastatic.com/images/photo/440314/1761252184/26087356455159750.jpg?scrop=1&amp;crop=1&amp;cpos=north&amp;w=110&amp;h=110" alt=""> <div data-v-5b109fc3="" class="num">1/4</div></div><div data-v-5b109fc3="" href="https://photo.zastatic.com/images/photo/440314/1761252184/6305611620118044.jpg" class="photoItem z-cursor-big"><img data-v-5b109fc3="" src="https://photo.zastatic.com/images/photo/440314/1761252184/6305611620118044.jpg?scrop=1&amp;crop=1&amp;cpos=north&amp;w=110&amp;h=110" alt=""> <div data-v-5b109fc3="" class="num">2/4</div></div><div data-v-5b109fc3="" href="https://photo.zastatic.com/images/photo/440314/1761252184/5832537473743391.jpg" class="photoItem z-cursor-big"><img data-v-5b109fc3="" src="https://photo.zastatic.com/images/photo/440314/1761252184/5832537473743391.jpg?scrop=1&amp;crop=1&amp;cpos=north&amp;w=110&amp;h=110" alt=""> <div data-v-5b109fc3="" class="num">3/4</div></div><div data-v-5b109fc3="" href="https://photo.zastatic.com/images/photo/440314/1761252184/25829281885635031.jpg" class="photoItem z-cursor-big"><img data-v-5b109fc3="" src="https://photo.zastatic.com/images/photo/440314/1761252184/25829281885635031.jpg?scrop=1&amp;crop=1&amp;cpos=north&amp;w=110&amp;h=110" alt=""> <div data-v-5b109fc3="" class="num">4/4</div></div></div>
`

const text3 = `
<div class="content"><table><tbody><tr><th><a href="http://album.zhenai.com/u/1761252184" target="_blank">普普通通</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>女士</td> <td><span class="grayL">居住地：</span>四川阿坝</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>46</td> <td><span class="grayL">学&nbsp;&nbsp;&nbsp;历：</span>高中及以下</td> <!----></tr> <tr><td width="180"><span class="grayL">婚况：</span>丧偶</td> <td width="180"><span class="grayL">身&nbsp;&nbsp;&nbsp;高：</span>160</td></tr></tbody></table> <div class="introduce">网络一线牵，相聚在空间，真情连四海，珍惜这段缘……(我还不是会员，所以回不了你们的信息，抱歉&lt;(_ _)&gt;)<br></div></div>
`

func main() {

	//	匹配字串
	//re := regexp.MustCompile("skyz319@gmail.com")
	//
	//match := re.FindString(text)
	//fmt.Println(match)

	//	正则匹配
	//re := regexp.MustCompile(`[a-zA-Z0-9._]+@[a-zA-Z0-9]+\.[a-zA-Z0-9.]+`)
	//match := re.FindAllString(text, -1)
	//fmt.Println(match)

	////	正则提取
	//re2 := regexp.MustCompile(`([a-zA-Z0-9._]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//match2 := re2.FindAllStringSubmatch(text, -1)
	//fmt.Println("match2:", match2)

	reg3 := regexp.MustCompile(`<div class="content">.*?<a href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>.*?<td width="180"><span class="grayL">性别：</span>([^<]+)</td>.*?</div>`)
	match3 := reg3.FindAllStringSubmatch(text3, -1)
	//fmt.Println("match3:", match3)

	for _, item := range match3 {
		str := strings.Replace(item[1], `\u002F`, "/", -1)

		fmt.Printf("%s ---- %s ---- %s\n", str, item[2], item[3])
	}
}
