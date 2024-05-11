package entity

import "time"

type UserInfo struct {
	Id              uint64     `json:"id"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	Username        string     `json:"username"`
	Password        string     `json:"password"`
	Avatar          string     `json:"avatar"`
	BackgroundImage string     `json:"background_image" gorm:"default:default_background.jpg"`
	Signature       string     `json:"signature"`
}

func (*UserInfo) TableName() string {
	return "userinfo"
}
