package unierr

var (
	Success       = NewErrCore(200, "success")
	ErrorParams   = NewErrCore(1, "参数错误")
	ErrorInternal = NewErrCore(2, "网络错误")
)

var (
	DifferentPassword     = NewErrCore(1001, "两次输入密码不一致")
	UserNameExist         = NewErrCore(1002, "用户名已存在")
	UsernameOrPasswordErr = NewErrCore(1003, "用户名或者密码错误")
	NoLogin               = NewErrCore(1004, "请登陆")
)
