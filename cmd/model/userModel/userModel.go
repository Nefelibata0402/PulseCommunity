package userModel

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
	Username        string `json:"username" form:"username" validate:"required"`
	Password        string `json:"password" form:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" validate:"required"`
}

func ValidateRegisterRequest(registerReq *RegisterRequest) error {
	validate := validator.New()
	return validate.Struct(registerReq)
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
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

func ValidateLoginRequest(loginReq *LoginRequest) error {
	validate := validator.New()
	return validate.Struct(loginReq)
}
