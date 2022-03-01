package service

import (
	"errors"

	"github.com/bingfenglai/gt/config"
	"github.com/bingfenglai/gt/conmon/constants"
	"github.com/bingfenglai/gt/conmon/helper"
	"github.com/bingfenglai/gt/model/entity"
	"github.com/bingfenglai/gt/model/shortcodegen"
	"go.uber.org/zap"

	"github.com/bingfenglai/gt/storage"
)

type IShortCodeService interface {

	// 创建短码并保存
	CreateShortCode(url string, isPerpetual, isMultiplex bool) (*entity.ShortCode, error)

	// 根据短码查找原链接
	FindLinkByCode(code string) (string, error)

	// 创建临时的短码
	createPerpetual(url string) (*entity.ShortCode, error)
}

type ShortCodeServiceImpl struct {
}

func (svc *ShortCodeServiceImpl) CreateShortCode(url string, isPerpetual, isMultiplex bool) (*entity.ShortCode, error) {
	zap.L().Info("收到入参",zap.Any("url",url))
	var shortcode *entity.ShortCode
	var err error
	urlMd5 := helper.ToMd5String32(url)
	
	
	if storage.ShortCodeStorage ==nil  {
		zap.L().Error("\n\nshort code store is null")
		return nil,err
	}
	
	if url == "" {
		return nil, errors.New("url不能为空")
	}

	// 创建临时短码
	if isPerpetual {

		return svc.createPerpetual(url)
		
	}

	// 先查询数据库中是否存在可复用的短码
	if isMultiplex {
		
		
		shortcode,err = storage.ShortCodeStorage.FindShortcodeByMd5(urlMd5)
		zap.L().Info("",zap.Any("sc: ",shortcode))
		if flag,_ :=helper.CheckErr(err);flag{
			return shortcode,nil
		}
		

	}


	if gen, err := shortcodegen.GetShortCodeGeneratorByMethod(shortcodegen.Md5Gen);err==nil{

		codes,errgen := gen.GenShortCode(url)

		if errgen!=nil {
			zap.L().Error(errgen.Error())
			return nil,errgen
		}
		

		if shortcode, err = entity.CreateShortCode(url,urlMd5,codes[0],constants.Registered_User_Code_Type);err!=nil{
			return nil,err
			
		}

		if flag,err :=storage.ShortCodeStorage.SaveOrUpdate(shortcode);!flag{
				return nil,err
		}


		return shortcode,nil

	}		



	return nil, err

}


func (svc *ShortCodeServiceImpl) FindLinkByCode(code string) (string, error){
	
	
	if code=="" {
		return "",errors.New("code不能为空")
	}

	if config.Conf.ShortCode.Length!=len(code) {
		return "",errors.New("code长度不正确")
		
	}

	if sc,err :=storage.ShortCodeStorage.FindOriginalUrlByShortCode(code);err!=nil{
		
		return "",err
	}else{
		return sc.Original,nil
	}

	
}


func (svc *ShortCodeServiceImpl) createPerpetual(url string) (*entity.ShortCode, error){



	return nil,nil
}
