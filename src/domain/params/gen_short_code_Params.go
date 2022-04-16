package params

import (
	"github.com/bingfenglai/gt/common/errors"
	"github.com/bingfenglai/gt/common/helper"
)

// 生成短码参数，游客创建的为临时code
type GenShortCodeParams struct {
	OriginalLink string `json:"original_link" binding:"required"`
	GroupId      int64  `json:"group_id"`
	IsMultiplex  bool   `json:"is_multiplex"`
}

func (param *GenShortCodeParams) Check() error {
	if param.OriginalLink == "" {
		return errors.New("原链接不能为空")
	}

	err := helper.CheckUrl(param.OriginalLink)

	if err != nil {
		return err
	}

	return nil
}
