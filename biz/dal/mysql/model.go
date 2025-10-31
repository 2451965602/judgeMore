package mysql

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserId    int64
	RoleId    string //实际上是我们业务过程中区分用户的主键
	UserName  string
	UserRole  string
	College   string
	Grade     string
	Major     string
	Email     string
	Status    int
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
