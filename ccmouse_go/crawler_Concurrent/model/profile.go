package model

import "encoding/json"

type Profile struct {
	Name     string   //	姓名
	Gender   string   //	性别
	UserInfo []string //	用户信息
	MateInfo []string //	择偶信息
	Photos   []string //	照片列表
}

func FromJsonObj(o interface{}) (Profile, error) {

	var profile Profile

	s, err := json.Marshal(o)
	if err != nil {

		return profile, err
	}

	err = json.Unmarshal(s, &profile)

	return profile, err
}
