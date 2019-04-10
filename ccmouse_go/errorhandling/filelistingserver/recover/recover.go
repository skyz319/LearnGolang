package main

import (
	"fmt"
)

func tryRecover() {

	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			//	捕获Panic中的错误
			fmt.Println("Error occurred:", err)
		} else {
			//	重新Panic
			panic(fmt.Sprintf("Don't know what to panic: %v", r))
		}
	}()

	//panic(errors.New("this is a error"))

	//a, b := 5, 0
	//a = a / b
	//fmt.Println(a)

	//	Panic一个未知错误
	panic(123)
}

func main() {

	tryRecover()
}
