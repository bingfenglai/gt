package helper

import "strings"

// 查找切片中是否包含指定元素
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {

		if len(val)==len(item) {
			if item == val {
				return i, true
			}
		}

		if strings.HasPrefix(val,item) {
			return i, true
		}
		
	}
	return -1, false
}
