package userData

type UserInfo struct {
	Id       int64
	UserId   uint64 `json:"user_id"`
	Username string
	Password string
}

func (*UserInfo) TableName() string {
	return "userinfo"
}
