package main

import (
	"fmt"
	"regexp"
)

const text = `My email is skyz319@gmail.com
email1 is cain.code@qq.com
email2 is 1691164@qq.com
sky.z@tom.com
skyz_319@tom.com
skyz_319@tom.com.cn
`

const text2 = `
<div class="logo f-fl" style="background-image:url(https://photo.zastatic.com/images/photo/308779/1235115550/4612939340901982.jpg?scrop=1&amp;crop=1&amp;cpos=north&amp;w=200&amp;h=200);" data-v-5b109fc3=""></div>`

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

	reg3 := regexp.MustCompile(`<div <div class="logo f-fl" style=".*? (^(\w+):\/\/([^/:]+)\/([^/:]+)\/(\d*)?\/(\d*)$[^?]+).*?" .*?></div>`)

	match3 := reg3.FindAllStringSubmatch(text2, -1)
	fmt.Println("match3:", match3)

}
