package seeder

import (
	"pos-v2-be/internal/enums"
	"pos-v2-be/internal/pkg/str"
	"pos-v2-be/internal/schema"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DefaultSetupSeeder(db *gorm.DB) {
	hash1, _ := str.HashPassword("admin123")

	datas := []*schema.User{
		{ID: uuid.New(), Email: "rizwar@gmail.com", Password: &hash1, Role: enums.RoleTypeADMIN},
	}

	db.Create(datas)
}
