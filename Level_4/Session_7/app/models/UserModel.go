package models

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	Id          string         `gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Username    string         `gorm:"index;unique;column:username;required"`
	Nickname    string         `gorm:"column:nickname;required"`
	Email       string         `gorm:"index;unique;column:email;required"`
	Password    string         `gorm:"column:password;required"`
	CreatedDate time.Time      `gorm:"column:created_date;autoCreateTime:true;not null"`
	UpdatedDate *time.Time     `gorm:"column:updated_date;autoUpdateTime:true"`
	DeletedDate gorm.DeletedAt `gorm:"index;column:deleted_date;softDelete:true"`
}

func (c *UserModel) TableName() string {
	return "users"
}
