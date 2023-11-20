package models

import (
	"database/sql"
	"time"

	"lin.com/gorm/utils"
)

type User struct {
	utils.GORM_MODEL
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}
