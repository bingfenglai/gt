package test

import (
	"fmt"
	"github.com/bingfenglai/gt/global"
	"github.com/bingfenglai/gt/model/entity"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"testing"
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

	return

}

func TestAddUser1(t *testing.T) {

	password, _ := bcrypt.GenerateFromPassword([]byte("gt@2022"), bcrypt.MinCost)

	fmt.Println(string(password))
}
