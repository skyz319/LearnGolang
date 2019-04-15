package main

import (
	"fmt"
	"time"
)

func chanDemo() {

	//	channel数组
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		//channels[i] = make(chan int)
		//go worker(i, channels[i])
		channels[i] = createWorker(i)
	}

	//	传入数据
	for i := 0; i < 10; i++ {

		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {

		channels[i] <- 'A' + i
	}

	//	避免main函数结束，休眠一毫秒以保证所有数据能进行打印
	time.Sleep(time.Millisecond)
}

func worker(id int, c chan int) {

	//	死循环，一直接收数据
	for {

		fmt.Printf("Worker %d received %d \n", id, <-c)
	}
}

//	返回channel
//	chan <- int代表是发送数据用的channel
//	<- chan int代表是收取数据用的channel
func createWorker(id int) chan<- int {

	c := make(chan int)

	go func() {

		//	死循环，一直接收数据
		for {

			fmt.Printf("Worker %d received %c \n", id, <-c)
		}

	}()

	return c
}

func bufferedChannel() {

	//	缓冲区3 在无接收者的情况下不会panic 超过缓冲区大小无接收者，则报错
	c := make(chan int, 3)
	go worker(0, c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)
}

func main() {

	//chanDemo()
	bufferedChannel()
}
