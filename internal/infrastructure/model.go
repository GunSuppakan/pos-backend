package infrastructure

import "gorm.io/gorm"

type Connections struct {
	DB *gorm.DB
}
