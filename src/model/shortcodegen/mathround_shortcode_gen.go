package shortcodegen

import (
	"errors"

	"math/rand"
	"strings"
	"time"

	"github.com/bingfenglai/gt/config"
	
)

// 数学函数生成，通过纳秒时间种子
// 对于随机数成成的code,可以采取预先生成的策略
type MathRoundShortCodeGenerator struct {
	genMethod string
}

func NewRoundShortCodeGenerator() *MathRoundShortCodeGenerator {

	gen := &MathRoundShortCodeGenerator{
		genMethod: MathRoundGen,
	}

	return gen
}

func (receiver *MathRoundShortCodeGenerator) GetGenMethod() string {
	return receiver.genMethod
}

func (receiver *MathRoundShortCodeGenerator) GenShortCode(link string) ([]string, error) {

	if link == "" {
		return nil, errors.New("参数link不能为空串")
	}

	return receiver.doGenShortCode()
}

func (receiver *MathRoundShortCodeGenerator) doGenShortCode() ([]string, error) {

	shortCodes := make([]string, config.Conf.ShortCode.Size)
	// 以当前纳秒数为种子
	rand.Seed(int64(time.Now().UnixNano()))
	max := len(chars()) - 1
	for i := 0; i <  config.Conf.ShortCode.Size; i++ {
		var shortCode []string
		for j := 0; j < config.Conf.ShortCode.Length; j++ {
			index := rand.Intn(max)
			//log.Default().Println("当前索引：",index)
			shortCode = append(shortCode, chars()[index])
		}
		shortCodes[i] = strings.Join(shortCode, "")
	}
	return shortCodes, nil
}
