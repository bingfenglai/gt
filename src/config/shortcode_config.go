package config

// 每批次生成的短码数最大值
const maxShortCodeGenSize = 6

type ShortCodeConfig struct {
	Length int
	Size int
}