package infra

import (
	"pos-v2-be/internal/initial/intf"

	"gorm.io/gorm"
)

type Infra struct {
	Upload     *Upload
	Service    *intf.Services
	Controller *intf.Controllers
	Repository *intf.Repositories
}

func NewInfra(db *gorm.DB) *Infra {
	// InitRedisInfra(db)
	upload := NewUpload()
	repositories := NewRepository(db)
	service := NewService(repositories, db)
	controller := NewController(service)

	return &Infra{
		Upload:     upload,
		Service:    service,
		Controller: controller,
		Repository: repositories,
	}
}
