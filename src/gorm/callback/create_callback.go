package callback

import (
	"github.com/bingfenglai/gt/oauth/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

func CreateWithTenantCallback(db *gorm.DB) {
	ctx := db.Statement.Context
	gCtx := ctx.(*gin.Context)
	var tenantId = ""
	var uid = -1
	if gCtx != nil {
		if user, err := utils.GetCurrentUser(gCtx.Request); err == nil {
			tenantId = user.TenantId
			uid, _ = strconv.Atoi(user.Uid)
			zap.L().Info("当前用户", zap.Any("user", user))
		}
	}

	if db.Error == nil && db.Statement.Error == nil {

		if tenantId != "" && tenantId != "-" {
			field := db.Statement.Schema.FieldsByName["TenantId"]

			if field != nil {
				zap.L().Info("field name is TenantId")
				tid, _ := strconv.Atoi(tenantId)
				err := field.Set(db.Statement.ReflectValue, tid)

				zap.Error(err)
			}
		}

		if uid != -1 {
			if createdByField := db.Statement.Schema.FieldsByName["CreatedBy"]; createdByField != nil {
				_ = createdByField.Set(db.Statement.ReflectValue, uid)
			}

		}

	}
}
