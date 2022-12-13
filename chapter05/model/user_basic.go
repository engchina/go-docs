package model

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Identity      string
	Name          string
	Password      string
	Phone         string
	Email         string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LogoutTime    time.Time
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
