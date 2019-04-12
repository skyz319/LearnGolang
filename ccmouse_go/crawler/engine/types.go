//	字段定义文件
package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult //	对下级页面的解析器
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

//	解决空parserFunc的问题
func NilParser([]byte) ParseResult {

	return ParseResult{}
}
