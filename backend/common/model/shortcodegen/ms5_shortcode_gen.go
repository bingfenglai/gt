package shortcodegen

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/bingfenglai/gt/config"
	
)

// 对链接计算md5生成code 在目前三种实现当中最慢
// 对于随机数成功的code,可以采取预先生成的策略，
type Md5ShortCodeGenerator struct {
	genMethod string
}

func NewMd5ShortCodeGenerator() *Md5ShortCodeGenerator {
	conver := &Md5ShortCodeGenerator{
		genMethod: Md5Gen,
	}

	return conver
}

func (receiver *Md5ShortCodeGenerator) GetGenMethod() string {
	return receiver.genMethod
}

func (receiver *Md5ShortCodeGenerator) GenShortCode(link string) ([]string, error) {

	if link == "" {
		return nil, errors.New("参数link不能为空")
	}

	defer func() {
		if err := recover(); err != nil {
			log.Default().Fatal(err)
		}
	}()

	return receiver.doGenShortCode(link)
}

func (receiver *Md5ShortCodeGenerator) doGenShortCode(link string) ([]string, error) {
	log.Default().Println("shortcode配置",config.Conf.ShortCode)
	// 用来存储生成的4组短链接
	codes := make([]string, config.Conf.ShortCode.Size)

	// 1.计算md5
	md5Byte := md5.New().Sum([]byte(link))
	md5Str := hex.EncodeToString(md5Byte[:])
	//log.Default().Println("生成的md5值：",md5Str)

	// 2.将md5值拆分为4组，每组8字节
	//log.Default().Println("md5长度： ",len(md5Str))
	max := len(chars()) - 1
	for i := 0; i <  config.Conf.ShortCode.Size; i++ {

		s := md5Str[i*8 : i*8+8]

		l, _ := strconv.ParseInt(s, 16, 32)
		// log.Default().Println("8字节的md5",l)
		//截取高位30位 与0x3fffffff(30位1)与操作, 即超过30位的忽略处理
		  l = 0x3FFFFFFF & l
		//   log.Default().Println("截取8字节的md5",l)
		//循环获得每组6位的字符串

		var shortCode []string
		move :=30/config.Conf.ShortCode.Length
		for i := 0; i < config.Conf.ShortCode.Length; i++ {
			//log.Default().Printf("l: %d",l)

			//两个位都为1时，结果才为1
			// index := 0x0000003D & l
			// 确保索引在字符集数组范围内
			index := int64(max) & l

			//log.Default().Printf("index： %d",index)
			//每5位的数字作为字母表的索引取得特定字符, 依次进行获得6位字符串
			shortCode = append(shortCode[:], chars()[index])

			// 各二进位全部右移若5位，对无符号数，高位补0
			l = l >> int64(move)

		}

		codes[i] = strings.Join(shortCode, "")
		// log.Default().Printf("md5 第%d组代码 %s",i+1,codes[i])

	}

	return codes, nil

}
