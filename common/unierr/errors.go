package unierr

var (
	Success       = NewErrCore(200, "success")
	ErrorParams   = NewErrCore(1, "参数错误")
	ErrorInternal = NewErrCore(2, "网络错误")
)

var (
	DifferentPassword                   = NewErrCore(1001, "两次输入密码不一致")
	UserNameExist                       = NewErrCore(1002, "用户名已存在")
	UsernameOrPasswordErr               = NewErrCore(1003, "用户名或者密码错误")
	NoLogin                             = NewErrCore(1004, "请登陆")
	UserNameOrPassword                  = NewErrCore(1005, "用户名或密码不能为空")
	UserNameOrPasswordOrConfirmPassword = NewErrCore(1006, "用户名或密码或确认密码不能为空")
)

var (
	ArticleTitleOrContentNotNil = NewErrCore(2001, "文章标题或内容不能为空")
	WithdrawArticleIDNotNIl     = NewErrCore(2002, "撤回文章Id不能为空")
)
