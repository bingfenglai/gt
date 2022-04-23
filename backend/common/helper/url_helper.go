package helper

import (
	"errors"
	"regexp"
)

func CheckUrl(url string) error {
	re := regexp.MustCompile(`(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`)
	result := re.FindAllStringSubmatch(url, -1)
	if result == nil {
		return errors.New("url不合法")
	}
	return nil
}
