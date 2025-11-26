package seeder

import (
	"gorm.io/gorm"
)

func RunSeeder(db *gorm.DB) {
	DefaultSetupSeeder(db)
}
