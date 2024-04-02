package errs

var (
	SuccessCode       = NewErrCore(0, "ok")
	InternalError     = NewErrCore(1, "Internal error")
	InvalidTokenError = NewErrCore(2, "Invalid JWT token")
	NoTokenError      = NewErrCore(3, "Missing JWT token")
	IllegalParams     = NewErrCore(4, "非法参数")
)

var (
	UsernameOrPasswordIsEmpty             = NewErrCore(1001, "用户名或密码为空")
	UsernameOrPasswordLenMore32Characters = NewErrCore(1002, "用户名或密码长度不能大于32个字符")
	UsernameOrPasswordFailToUpload        = NewErrCore(1003, "用户名或密码上传失败")
	UsernameExist                         = NewErrCore(1004, "该用户名已存在")
	UsernameNotExist                      = NewErrCore(1005, "用户名不存在")
	PasswordWrong                         = NewErrCore(1006, "密码错误")
	UserNotExist                          = NewErrCore(1007, "用户不存在")
	TokenNotExist                         = NewErrCore(1008, "Token不存在")
	UserIdInvalid                         = NewErrCore(1009, "user_id不合法")
)
