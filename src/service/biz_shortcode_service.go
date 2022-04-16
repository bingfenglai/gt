package service

import (
	"context"
	"errors"
	"github.com/bingfenglai/gt/domain/params"
	"github.com/bingfenglai/gt/oauth/utils"
	"gorm.io/gorm"
	"time"

	"github.com/bingfenglai/gt/common/constants"
	"github.com/bingfenglai/gt/common/helper"
	"github.com/bingfenglai/gt/common/model/shortcodegen"
	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/domain/entity"

	"go.uber.org/zap"

	"github.com/bingfenglai/gt/storage"
)

const SHORTCODE_CACHE_PREFIX = "sc:"

type IShortCodeService interface {

	// 创建短码并保存
	CreateShortCode(url string, isPerpetual, isMultiplex bool) (*entity.ShortCode, error)

	CreateShortCodeWithContext(params params.GenShortCodeParams, ctx context.Context) (*entity.ShortCode, error)
	// 根据短码查找原链接
	FindLinkByCode(code string) (*entity.ShortCode, error)

	// 创建临时的短码
	createPerpetual(url string) (*entity.ShortCode, error)
}

type shortCodeServiceImpl struct {
}

func (svc *shortCodeServiceImpl) CreateShortCodeWithContext(params params.GenShortCodeParams, ctx context.Context) (sc *entity.ShortCode, err error) {

	if err = params.Check(); err != nil {
		return
	}

	if _, err := utils.GetCurrentUIdWithContext(ctx); err != nil {
		return svc.createPerpetual(params.OriginalLink)
	}

	urlMd5 := helper.ToMd5String32(params.OriginalLink)

	// 先查询数据库中是否存在可复用的短码
	if params.IsMultiplex {

		sc, err = storage.ShortCodeStorage.FindShortcodeByMd5(urlMd5)
		zap.L().Info("", zap.Any("sc: ", sc))
		if err == nil {
			return
		}

	}

	code, err := svc.genShortCode(params.OriginalLink, shortcodegen.Md5Gen)

	if err != nil {
		return nil, err
	}

	if sc, err = entity.CreateShortCode(params.OriginalLink, urlMd5, code, constants.Registered_User_Code_Type); err != nil {
		return nil, err
	}

	if flag, err := storage.ShortCodeStorage.SaveOrUpdate(ctx, sc); !flag {
		return nil, err
	}

	return

}

func (svc *shortCodeServiceImpl) CreateShortCode(url string, isPerpetual, isMultiplex bool) (*entity.ShortCode, error) {
	zap.L().Info("收到入参", zap.Any("url", url))
	var shortcode *entity.ShortCode
	var err error
	urlMd5 := helper.ToMd5String32(url)

	if url == "" {
		return nil, errors.New("url不能为空")
	}

	// 创建临时短码
	if isPerpetual {

		return svc.createPerpetual(url)

	}

	// 先查询数据库中是否存在可复用的短码
	if isMultiplex {

		shortcode, err = storage.ShortCodeStorage.FindShortcodeByMd5(urlMd5)
		zap.L().Info("", zap.Any("sc: ", shortcode))
		if flag, _ := helper.CheckErr(err); flag {
			return shortcode, nil
		}

	}

	if gen, err := shortcodegen.GetShortCodeGeneratorByMethod(shortcodegen.Md5Gen); err == nil {

		codes, errgen := gen.GenShortCode(url)

		if errgen != nil {
			zap.L().Error(errgen.Error())
			return nil, errgen
		}

		if shortcode, err = entity.CreateShortCode(url, urlMd5, codes[0], constants.Registered_User_Code_Type); err != nil {
			return nil, err

		}

		if flag, err := storage.ShortCodeStorage.SaveOrUpdate(nil, shortcode); !flag {
			return nil, err
		}

		return shortcode, nil

	}

	return nil, err

}

func (svc *shortCodeServiceImpl) FindLinkByCode(code string) (sc *entity.ShortCode, err error) {

	if code == "" {
		return nil, errors.New("code不能为空")
	}

	if config.Conf.ShortCode.Length != len(code) {
		return nil, errors.New("code长度不正确")

	}

	if err = CacheService.Get(SHORTCODE_CACHE_PREFIX+code, sc); err != nil {
		sc, err = storage.ShortCodeStorage.FindOriginalUrlByShortCode(code)
	}

	return

}

func (svc *shortCodeServiceImpl) createPerpetual(url string) (sc *entity.ShortCode, err error) {
	var code = ""
	if code, err = svc.genShortCode(url, shortcodegen.MathRoundGen); err == nil {
		urlMd5 := helper.ToMd5String32(url)
		sc, err = entity.CreateShortCode(url, urlMd5, code, constants.Anonymous_Code_Type)
	}

	if sc != nil {
		CacheService.Set(SHORTCODE_CACHE_PREFIX+code, sc, time.Minute*30)
	}

	return
}

func (svc shortCodeServiceImpl) genShortCode(url, genMethod string) (code string, err error) {

	if gen, err := shortcodegen.GetShortCodeGeneratorByMethod(genMethod); err == nil {

		codes, err := gen.GenShortCode(url)

		if err != nil {
			zap.L().Error(err.Error())
			return
		}
		var code = ""
		for _, c := range codes {
			if _, err := svc.FindLinkByCode(code); err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					code = c
				}
			}
		}

		if code == "" {
			return svc.genShortCode(url, shortcodegen.MathRoundGen)
		}

	}

	return
}
