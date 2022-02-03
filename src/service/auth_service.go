package service

import (
	"github.com/bingfenglai/gt/pojo/params"
	"github.com/bingfenglai/gt/pojo/response"
)

// 认证授权服务接口
type IAuthService interface {
	Authentication(params params.AuthcParams) (*response.TokenResponse, error)
}
