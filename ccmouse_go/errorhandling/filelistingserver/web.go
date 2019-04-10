package main

import (
	"LearnGolang/ccmouse_go/errorhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

//	用户错误 接口
type userError interface {
	error
	Message() string
}

type appHandler func(write http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {

	return func(write http.ResponseWriter, request *http.Request) {

		//	recover
		defer func() {

			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(write, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		//	normal error handle
		err := handler(write, request)

		if err != nil {

			log.Printf("Error handling request: %s", err.Error())

			//	处理错误信息 区分用户可见错误
			if userErr, ok := err.(userError); ok {

				http.Error(write, userErr.Message(), http.StatusBadRequest)

				return
			}

			code := http.StatusOK

			switch {

			case os.IsNotExist(err): //	文件不存在
				log.Printf("no such file or directory")
				code = http.StatusNotFound

			case os.IsPermission(err): // 没有权限
				log.Printf("Permission error")
				code = http.StatusForbidden

			default:
				log.Println("default error")
				code = http.StatusInternalServerError
			}

			http.Error(write, http.StatusText(code), code)

		}
	}
}

func main() {

	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
