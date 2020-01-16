package context

import(
	"github.com/ahmad2smile/battle-gin/models"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func Setup(db *gorm.DB) {
	Db = db
	db.LogMode(true)
	db.AutoMigrate(&models.Game{})
}
