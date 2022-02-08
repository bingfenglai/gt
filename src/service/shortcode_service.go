package service

import (
	"errors"
)

type IShortCodeService interface {

	// 创建短码并保存
	CreateShortCodeAndSave(url string, isPerpetual, isMultiplex bool) (bool, error)

	// 根据短码查找原链接
	FindLinkByCode(code string) (string, error)
}

type ShortCodeService struct {
}

func (this *ShortCodeService) CreateShortCodeAndSave(url string, isPerpetual, isMultiplex bool) (bool, error) {

	if url == "" {
		return false, errors.New("url不能为空")
	}

	// 创建临时短码
	if !isPerpetual {

	}

	// 先查询数据库中是否存在可复用的短码
	if isMultiplex {
		//urlMd5 := helper.ToMd5String32(url)
		//
		//shortCode := entity.ShortCode{}
		//
		//global.DB.f
	}

	return true, nil

}
