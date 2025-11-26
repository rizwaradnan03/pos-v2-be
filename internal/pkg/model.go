package pkg

import "pos-v2-be/internal/schema"

func ModelsToMigrate() []interface{} {

	return []interface{}{
		&schema.User{},
		&schema.Log{},
	}
}
