package string_util

import (
	"strings"
	"unicode/utf8"
)

/**
去除首尾空格
*/
func TrimBlank(value string) string {
	return strings.Trim(value, " ")
}

func TrimUtf8BOM(value string) string {
	bytes := []byte(value)
	count := len(bytes)
	if count <= 3 {
		return value
	}
	return string(bytes[3:count])
}

func IsUtf8NotBOM(value string) bool {
	if utf8.ValidString(value) {
		bytes := []byte(value)
		count := len(bytes)
		if count < 3 {
			return true
		} else {
			return bytes[0] != 0xef || bytes[1] != 0xbb || bytes[2] != 0xbf
		}
	}
	return false
}

func IsUtf8BOM(value string) bool {
	if utf8.ValidString(value) {
		bytes := []byte(value)
		count := len(bytes)
		if count < 3 {
			return false
		} else {
			result := bytes[0] == 0xef && bytes[1] == 0xbb && bytes[2] == 0xbf
			return result
		}
	}
	return false
}
