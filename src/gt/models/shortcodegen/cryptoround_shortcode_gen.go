package shortcodegen

import (
	"crypto/rand"
	"errors"

	"math/big"
	"strings"
)

// 种子跟 进程数+内存占用长度+时间戳 有关
// 比数学随机函数慢10倍
// 对于随机数成成的code,可以采取预先生成的策略，
type CryptRoundShortCodeGenerator struct {
	GenMethod string
}

func (receiver *CryptRoundShortCodeGenerator) GetGenMethod() string {
	return receiver.GenMethod
}

func NewCryptRoundShortCodeGenerator() *CryptRoundShortCodeGenerator {

	return &CryptRoundShortCodeGenerator{
		GenMethod: "cryptRound",
	}
}

func (receiver *CryptRoundShortCodeGenerator) GenShortCode(link string) ([]string, error) {
	if link == "" {
		return nil, errors.New("参数link不能为空串")
	}

	return receiver.doGenShortCode()

}

func (receiver *CryptRoundShortCodeGenerator) doGenShortCode() ([]string, error) {
	shortCodes := make([]string, 4)

	max := len(chars())
	for i := 0; i < 4; i++ {
		var shortCode []string
		for j := 0; j < 6; j++ {
			index, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
			//log.Default().Println("当前索引：",index.Int64())

			shortCode = append(shortCode, chars()[index.Int64()])
		}

		shortCodes[i] = strings.Join(shortCode, "")

	}

	return shortCodes, nil
}
