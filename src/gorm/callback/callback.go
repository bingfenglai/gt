package callback

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
)

func CreatedCallback(db *gorm.DB)  {

	ctx := context.Background()
	t := ctx.Value("TenantId")
	
	log.Default().Println(ctx)
	zap.L().Info("ctx",zap.Any("",t))

	if db.Error == nil && db.Statement.Error == nil {
		for name, _ := range db.Statement.Schema.FieldsByName {
			//zap.L().Info(field.Name)
			if name=="TenantId" {
				zap.L().Info("field name is TenantId")
			}
		}
	}


}