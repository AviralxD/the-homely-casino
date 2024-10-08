// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
)

type Bet struct {
	Name   string
	Amount int32
}

type Player struct {
	Name     string
	Totalbet sql.NullInt32
	Pack     sql.NullBool
	Show     sql.NullBool
}

type Round struct {
	Currentbet int32
	Totalbet   int32
	Playernum  sql.NullInt32
}
