package test

import (
	"fmt"
	"github.com/bingfenglai/gt/domain/entity"
	"github.com/bingfenglai/gt/global"

	"log"
	"testing"
	"time"

	"github.com/bingfenglai/gt/service"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func TestAddUser(t *testing.T) {

	password, _ := bcrypt.GenerateFromPassword([]byte("gt@2022"), bcrypt.MinCost)

	zap.L().Info("密码", zap.String("password", string(password)))
	//context.WithValue(nil,"t","tttttttttttttt")
	u := entity.User{
		Username:  "aliyun_969392",
		Password:  string(password),
		Email:     "bingfengdev1@aliyun.com",
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
	service.UserService.FindUserByUId(1)
	end := time.Now().Nanosecond()

	log.Default().Println("不使用缓存", end-start)

	start1 := time.Now().Nanosecond()
	service.UserService.FindUserByUIdWithCache(1)
	end1 := time.Now().Nanosecond()

	log.Default().Println("使用缓存1", end1-start1)

	start2 := time.Now().Nanosecond()
	service.UserService.FindUserByUIdWithCache(1)
	end2 := time.Now().Nanosecond()

	log.Default().Println("使用缓存2", end2-start2)

}


func TestFindUser(t *testing.T){
	user,err :=service.UserService.FindUserByEmail("bingfenglai.dev@gmail.com")

	if err!=nil {
		t.Error(err)
	}else{
		log.Default().Println(user.Username)
	}


}
