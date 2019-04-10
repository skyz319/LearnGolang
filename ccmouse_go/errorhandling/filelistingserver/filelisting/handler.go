package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

//	用户错误接口的实现
func (e userError) Error() string {
	return e.Message()
}

//	用户错误接口的实现
func (e userError) Message() string {

	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {

	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start with " + prefix)
	}

	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {

		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {

		return err
	}

	writer.Write(all)

	return nil
}
