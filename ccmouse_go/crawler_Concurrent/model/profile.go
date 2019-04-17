package model

type Profile struct {
	UserURL  string   //	用户网页
	Name     string   //	姓名
	Gender   string   //	性别
	UserInfo []string //	用户信息
	MateInfo []string //	择偶信息
	Photos   []string //	照片列表
}
