package main

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime:milli"`
	DeleteAt     gorm.DeletedAt `gorm:"index"`
}
