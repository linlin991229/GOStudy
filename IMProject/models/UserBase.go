package models

import (
	"lin.com/im/utils"
)

type UserBase struct {
	utils.GORM_MODEL
	Username      string `json:"username" example:"lin"`
	Password      string `json:"password" example:"123456"`
	Email         string `json:"email" example:"lin@qq.com"`
	Phone         string `json:"phone" example:"18989190946"`
	Identity      string `json:"identity"`
	ClientId      string `json:"client_id"`
	ClientPort    string `json:"client_port"`
	LoginTime     uint64 `json:"login_time"`
	LogOutTime    uint64 `json:"logout_time"`
	IsLogout      bool   `json:"is_logout"`
	HeartbeatTime uint64 `json:"heartbeat_time"`
	DeviceInfo    string `json:"device_info"`
}
