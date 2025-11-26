package initial

import (
	"pos-v2-be/internal/initial/infra"
	"pos-v2-be/internal/seeder"

	"gorm.io/gorm"
)

type Init struct {
	Infra *infra.Infra
}

func NewInit(db *gorm.DB) *Init {
	infra := infra.NewInfra(db)
	seeder.RunSeeder(db)

	return &Init{
		Infra: infra,
	}
}
