package entity

import (
	"github.com/bingfenglai/gt/global"
	"github.com/jinzhu/gorm"
)

type Client struct {
	gorm.Model
	// 客户端业务主键
	ClientBizId string `gorm:"not null;size:32"`
	Domain      string `gorm:"not null;size:128"`
	Secret      string `gorm:"not null;size:128"`
	CreatedBy   int64
	UpdatedBy   int64
	Status      int    `gorm:"default:0"`
	Remark      string `gorm:"not null;size:64"`
}

func (c *Client) TableName() string {
	return "tb_sys_client"
}

func (c *Client) GetByBizId(id string) error {
	err := global.DB.Where("client_biz_id = ?", id).Take(&c).Error

	return err

}

func (c *Client) Create() (bool, error) {

	err := global.DB.Begin().Create(&c).Commit().Error

	return err == nil, err
}
