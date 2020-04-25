package utils

import (
	"encoding/json"
)

//JSONToStr 序列化json对象为字符串
func JSONToStr(v interface{}) string {
	content, err := json.Marshal(v)
	if err != nil {
		Println(err)
		return ""
	}
	return Bytes2Str(content)
}

//JSONFromStr 解析字符串为json对象
func JSONFromStr(s string, v interface{}) {
	err := json.Unmarshal(Str2Bytes(s), v)
	if err != nil {
		Println(err)
	}
}
