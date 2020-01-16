package models

import "database/sql/driver"

type PlayerRole string

const (
	Attacker PlayerRole = "Attacker"
	Defender PlayerRole = "Defender"
)

func (playerRole *PlayerRole) Scan(value interface{}) error {
	*playerRole = PlayerRole(value.([]byte))

	return nil
}

func (playerRole PlayerRole) Value() (driver.Value, error) {
	return string(playerRole), nil
}
