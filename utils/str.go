package utils

import "strings"
import "strconv"

//IsBlankStr  判断字符串是否为空
func IsBlankStr(str1 string) bool {
	if len(strings.TrimSpace(str1)) == 0 {
		return true
	}
	return false
}

//IsBlankStr1 判断字符串是否为空
func IsBlankStr1(str *string) bool {
	if nil == str || len(strings.TrimSpace(*str)) == 0 {
		return true
	}
	return false
}

//Str2Bytes 字符串转byte数组
func Str2Bytes(str string) []byte {
	return []byte(str)
}

//Bytes2Str byte数组转字符串
func Bytes2Str(bytes []byte) string {
	return string(bytes)
}

//Str2Int 字符串转int
func Str2Int(value string) (int, error) {
	return strconv.Atoi(value)
}

//Int2Str int转字符串
func Int2Str(value int) string {
	return string(value)
}
