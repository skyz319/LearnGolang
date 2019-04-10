package main

// 斐波那契数列
import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type intGen func() int

func fibonacci() intGen {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return a
	}
}

func (g intGen) Read(p []byte) (n int, err error) {

	next := g()
	//	斐波那契数列数字大于指定时，中止
	if next > 100000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	return strings.NewReader(s).Read(p)
}

//	为函数实现接口
func printFileContents(reader io.Reader) {

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {

	f := fibonacci()

	//for i := 1; i < 10; i++ {
	//	fmt.Println(f())
	//}

	printFileContents(f)
}
