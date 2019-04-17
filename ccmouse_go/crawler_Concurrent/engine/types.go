//	字段定义文件
package engine

//	排除模式 排除指定性别 为空不排除
const Exclusion = ""

//	限制抓取数量
const LimitCity = false
const CityNum = 20
const LimitUser = false
const UserNum = 10

//	ElasticSearch 相关
const DataBaseName = "dating_profile"
const TableName = "zhenai"

type Item struct {
	//	公用数据URL和ID
	Url  string
	Id   string
	Type string //	项目的表名

	//	适配不同项目使用的数据体
	Payload interface{}
}

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult //	对下级页面的解析器
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

//	解决空parserFunc的问题
func NilParser([]byte) ParseResult {

	return ParseResult{}
}
