package service

import (
	"errors"
	"time"

	"github.com/bingfenglai/gt/conmon/helper"
	"github.com/bingfenglai/gt/model/entity"
	"github.com/bingfenglai/gt/storage"
)

type IShortCodeLogService interface {
	// 保存
	Save(shortcodeLog *entity.ShortcodeLog)(bool,error)

	Create(shorCodeId uint64,userAgent string,ip string) (bool,error)
}

type ShortCodeLogServiceImpl struct {
	
}

func (svc *ShortCodeLogServiceImpl) Save(shortcodeLog *entity.ShortcodeLog)(bool,error){

	if shortcodeLog==nil {
		return false,errors.New("参数不能为空")
	}

	return storage.ShortCodeLogStorage.SaveOrUpdate(shortcodeLog)

}


func (svc *ShortCodeLogServiceImpl) Create(shorCodeId uint64,userAgent string,ip string) (bool,error){
	us :=helper.ParseUserAgent(userAgent)
	return svc.Save(&entity.ShortcodeLog{
		ShortcodeId: shorCodeId,
		UserAgent: userAgent,
		Ip: uint32(helper.Ip2Long(ip)),
		OperationSystem: us.OS,
		Client:  us.Name,
		AccessTime: time.Now(),
	})
}
