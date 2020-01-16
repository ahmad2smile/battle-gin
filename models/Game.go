package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type GameDTO struct {
	PlayerRole PlayerRole `json:"playerRole"`
}

type Game struct {
	ID         uuid.UUID
	PlayerRole PlayerRole
}

// trigger not working for MYSQL
func (game *Game) BeforeSave(scope *gorm.Scope) error {
	newId := uuid.NewV4()

	scope.SetColumn("id", newId)

	return nil
}
