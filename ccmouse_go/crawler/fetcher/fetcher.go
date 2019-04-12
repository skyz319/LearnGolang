//	抓取器
package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return nil, fmt.Errorf("status code is error! status code :%d\n", resp.StatusCode)
	}

	/*
		读取Body内容
	*/
	//	获取字符集
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	//	转化为utf-8
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)

}

/*
	识别HTML的charset
*/
func determineEncoding(r *bufio.Reader) encoding.Encoding {

	//	读取前1024字节
	bytes, err := r.Peek(1024)
	if err != nil {

		log.Printf("Fetcher error: %v", err)
		//	默认UTF-8
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
