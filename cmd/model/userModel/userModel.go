package userModel

import "errors"

type RegisterRequest struct {
	UserName        string `json:"userName" form:"userName"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword"`
}

func (r RegisterRequest) Verify() error {
	if !r.VerifyPassword() {
		return errors.New("两次输入密码不一致")
	}
	return nil
}

func (r RegisterRequest) VerifyPassword() bool {
	return r.Password == r.ConfirmPassword
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
