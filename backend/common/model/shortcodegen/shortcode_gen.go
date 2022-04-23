package shortcodegen

import (
	"errors"

	"github.com/bingfenglai/gt/common/constants"
)

var genMap map[string]ShortCodeGenerator

const (
	Md5Gen         = "Md5Gen"
	MathRoundGen   = "MathRoundGen"
	CryptoRoundGen = "CryptoRoundGen"
)

// 链接长短转换接口
type ShortCodeGenerator interface {

	// 将长链接转换为4组code
	GenShortCode(link string) ([]string, error)

	// 获取当前生成短码的方式
	GetGenMethod() string
}

// 根据成生短码的方式
func GetShortCodeGeneratorByMethod(genMethod string) (ShortCodeGenerator, error) {

	if v, ok := genMap[genMethod]; ok {

		return v, nil
	} else {
		return nil, errors.New("未找到对应的短码生成器")
	}
}

// 将生成器加载到map中
func init() {
	genMap = make(map[string]ShortCodeGenerator)

	md5ShortCodeGen := NewMd5ShortCodeGenerator()
	genMap[md5ShortCodeGen.genMethod] = md5ShortCodeGen

	mathRoundGen := NewRoundShortCodeGenerator()
	genMap[mathRoundGen.genMethod] = mathRoundGen

	cryptoRoundGen := NewCryptRoundShortCodeGenerator()
	genMap[cryptoRoundGen.genMethod] = cryptoRoundGen

}

// const chars := []string{"q","w","e","r","t","y","u","i","o","p",
// "a","s","d","f","g","h","j","k","l","z",
// "x","c","v","b","n","m","1","2","3","4",
// "5","6","7","8","9","0","Q","W","E","R",
// "T","Y","U","I","O","P","A","S","D","F",
// "G","H","J","K","L","Z","X","C","V","B",
// "N","M"}

// code 待选字符
func chars() []string {

	return constants.Chars()
}
