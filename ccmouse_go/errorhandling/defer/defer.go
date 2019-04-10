package main

import (
	"LearnGolang/ccmouse_go/functional/fib"
	"bufio"
	"fmt"
	"os"
)

//	defer 相当于是栈，先进后出。后加入的defer 将在函数结束时先调用
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func writeFile(filename string) {

	//	新建文件
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	//	将内容写进缓存
	writer := bufio.NewWriter(file)
	//	将缓存内容写入文件
	defer writer.Flush()

	f := fib.Fibonacci()

	//	写入前20个
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {

	//tryDefer()
	writeFile("fib.txt")
}
