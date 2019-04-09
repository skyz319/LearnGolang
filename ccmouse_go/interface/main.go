package main

import (
	"LearnGolang/ccmouse_go/interface/real"
	"fmt"
)

//	接口
type Retriever interface {
	Get(url string) string
}

//	使用者
func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {

	var r Retriever
	//	实现者
	r = real.Retriever{}
	fmt.Println(download(r))
	fmt.Println()
	fmt.Printf("%T %v", r, r)
}
