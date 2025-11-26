package schema

import (
	"pos-v2-be/internal/enums"
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Type        string    `json:"type"`
	Description string    `json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Username     string         `gorm:"unique" json:"username"`
	Password     *string        `gorm:"size:255" json:"-"`
	Role         enums.RoleType `gorm:"type:varchar(20);default:'AUTHOR'" json:"role"`
	RefreshToken *string        `json:"refresh_token"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// CommissionBatch     []*CommissionBatch   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"commission_batch,omitempty"`
}

type Ingredient struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name string    `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	IngredientParticular     []*IngredientParticular   `gorm:"foreignKey:IngredientID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"ingredient,omitempty"`
}

type IngredientParticular struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	IngredientID uuid.UUID `json:"ingredient_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
