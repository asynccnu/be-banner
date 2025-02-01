package dao

import (
	"github.com/asynccnu/be-banner/domain"
	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Banner{})
}
