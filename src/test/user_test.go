package test

import (
	"fmt"
	"github.com/bingfenglai/gt/global"
	"github.com/bingfenglai/gt/model/entity"
	"github.com/bingfenglai/gt/service"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
	"time"
)

func TestAddUser(t *testing.T) {

	password, _ := bcrypt.GenerateFromPassword([]byte("gt@2022"), bcrypt.MinCost)

	zap.L().Info("密码", zap.String("password", string(password)))

	u := entity.User{
		Username:  "969391",
		Password:  string(password),
		Email:     "bingfenglai.dev@gmail.com",
		CreatedBy: 0,
		UpdatedBy: 0,
		Status:    0,
		TenantId:  0,
	}
	if err := global.DB.Begin().Save(&u).Commit().Error; err != nil {
		zap.L().Error(err.Error())
		return
	}

	

}

func TestAddUser1(t *testing.T) {

	password, _ := bcrypt.GenerateFromPassword([]byte("gt@2022"), bcrypt.MinCost)

	fmt.Println(string(password))
}

func TestCache(t *testing.T) {

	start := time.Now().Nanosecond()
	service.UserService.FindUserByUsername("969391")
	end := time.Now().Nanosecond()

	log.Default().Println("不使用缓存",end-start)

	start1 := time.Now().Nanosecond()
	service.UserService.FindUserByUsernameWithCache("969391")
	end1 := time.Now().Nanosecond()

	log.Default().Println("使用缓存1",end1-start1)

	start2 := time.Now().Nanosecond()
	service.UserService.FindUserByUsernameWithCache("969391")
	end2 := time.Now().Nanosecond()

	log.Default().Println("使用缓存2",end2-start2)






}
