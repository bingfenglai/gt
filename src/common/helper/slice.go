package helper

import (
	"strings"
)

// 查找切片中是否包含指定元素
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {

		if len(val) == len(item) {
			if item == val {
				return i, true
			}
		}

		if strings.HasPrefix(val, item)&&item[len(item)+1:]=="*" {
			return i, true
		}

	}
	return -1, false
}

func Match(og, tg string) (origin, target string) {

	if og == tg {
		return og, tg
	}

	// *字符匹配
	if strings.Contains(og, "*") {
		index := strings.Index(og, "*")
		if index+1 > len(og) {
			og = og[:index-1]
		} else {
			og = og[:index-1] + og[index+1:]
		}

		if index < len(tg) {
			temp1 := tg[:index-1]
			temp2 := tg[index:]
			if strings.Contains(temp2, "/") {
				index1 := strings.Index(temp2, "/")
				tg = temp1 + temp2[index1:]
			} else {
				tg = temp1
			}
		}

	}

	if strings.Contains(og, "*") {
		return Match(og, tg)
	}

	return og, tg
}
