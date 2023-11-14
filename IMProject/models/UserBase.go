package models

type UserBase struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Identity      string `json:"identity"`
	ClientId      string `json:"client_id"`
	ClientPort    string `json:"client_port"`
	LoginTime     uint64 `json:"login_time"`
	LogOutTime    uint64 `json:"logout_time"`
	IsLogOut      bool   `json:"is_logout"`
	HeartbeatTime uint64 `json:"heartbeat_time"`
	DeviceInfo    string `json:"device_info"`
}
