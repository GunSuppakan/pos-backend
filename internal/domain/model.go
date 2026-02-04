package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Model struct
type Model struct {
	Seq       int64      `json:"seq" gorm:"primary_key;auto_increment:true;"`
	Uid       uuid.UUID  `json:"uid" gorm:"primary_key;index"`
	CreatedAt time.Time  `json:"created_at" gorm:"index"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

type ModelV2 struct {
	Seq       int64      `json:"seq" gorm:"autoIncrement;uniqueIndex"`
	Uid       uuid.UUID  `json:"uid" gorm:"primary_key;index"`
	CreatedAt time.Time  `json:"created_at" gorm:"index"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

// BeforeCreate hook table
func (m *Model) BeforeCreate(tx *gorm.DB) error {
	if uuid.Equal(m.Uid, uuid.Nil) {
		m.Uid = uuid.NewV4()
	}
	m.Seq = time.Now().UnixNano()
	return nil
}
