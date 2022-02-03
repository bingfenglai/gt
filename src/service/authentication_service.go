package service

import (
	"errors"
	"github.com/bingfenglai/gt/conmon/helper"
	"github.com/bingfenglai/gt/pojo/params"
	"github.com/bingfenglai/gt/pojo/response"
	"github.com/hako/branca"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

var upAuthcService IAuthService

var refreshTokenAuthcService IAuthService

var br branca.Branca

type AuthenticationService struct {
}

type usernamePasswordAuthenticationService struct {
}

type refreshTokenAuthenticationService struct {
}

func (authc *AuthenticationService) Authentication(params params.AuthcParams) (*response.TokenResponse, error) {

	if params.GrantType == "password" {
		return upAuthcService.Authentication(params)
	}

	if params.GrantType == "refreshToken" {
		return refreshTokenAuthcService.Authentication(params)
	}

	return nil, errors.New("暂不支持该认证方式")
}

func (authc *usernamePasswordAuthenticationService) Authentication(params params.AuthcParams) (*response.TokenResponse, error) {
	userDto, err := UserService.FindUserByUsername(params.Principal)

	if err != nil {
		return nil, err
	}

	check, _ := PasswordEncodeService.check(params.Certificate, userDto.Password)

	if !check {
		return nil, errors.New("用户名或密码错误")
	}

	md5String := helper.ToMd5String32("gt.com")
	zap.L().Info(strconv.Itoa(len(md5String)))
	newBranca := branca.NewBranca(md5String)
	newBranca.SetTTL(2700)

	accessToken, err := newBranca.EncodeToString(params.Principal)

	if err != nil {
		zap.L().Error(err.Error())
	}

	accessToken = "Bearer " + accessToken

	refreshToken, err := newBranca.EncodeToString(userDto.Username + "&" + params.Certificate)

	if err != nil {
		zap.L().Error(err.Error())
	}

	newBranca.SetTTL(2700 * 2)

	refreshToken = "Bearer " + refreshToken
	token := response.TokenResponse{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		ExpiresIn:        2700,
		TokenType:        "Bearer",
		RefreshExpiresIn: 5400,
	}

	return &token, nil
}

func (authc *refreshTokenAuthenticationService) Authentication(p params.AuthcParams) (*response.TokenResponse, error) {

	s := strings.Split(p.Principal, "Bcrypt ")[1]

	md5String := helper.ToMd5String32("gt.com")
	zap.L().Info(strconv.Itoa(len(md5String)))
	newBranca := branca.NewBranca(md5String)

	s, err := newBranca.DecodeToString(s)

	if err != nil {
		return nil, err
	}

	usernamePassword := strings.Split(s, "&")

	return upAuthcService.Authentication(params.AuthcParams{
		Principal:   usernamePassword[0],
		Certificate: usernamePassword[1],
		GrantType:   "password",
	})

}

func init() {
	upAuthcService = &usernamePasswordAuthenticationService{}
	refreshTokenAuthcService = &refreshTokenAuthenticationService{}
}
