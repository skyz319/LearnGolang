//	字段定义文件
package engine

//	排除模式 排除指定性别 为空不排除
const Exclusion = ""

//	限制抓取数量
const Limit = true
const CityNum = 5
const UserNum = 10

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
