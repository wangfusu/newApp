package errcode

var (
	ErrorGetUserListFail   = NewError(20010001, "获取用户列表失败")
	ErrorCreateUserFail    = NewError(20010002, "创建用户失败")
	ErrorUpdateUserFail    = NewError(20010003, "更新用户失败")
	ErrorDeleteUserFail    = NewError(20010004, "删除用户失败")
	ErrorCountUserFail     = NewError(20010005, "统计用户失败")
	ErrorCheckUserFail     = NewError(20010006, "user名已被注册")
	ErrorCheckEmailFail    = NewError(20010007, "email已被注册")
	ErrorGetEmailCodeFail  = NewError(20010008, "获取邮箱验证码失败")
	ErrorEmailCodeNonValid = NewError(20010009, "邮箱验证码失效")
	ErrorEmailCodeFail     = NewError(20010010, "邮箱验证码不匹配")
	ErrorNotFineEmailCode  = NewError(20010011, "邮箱验证码参数不存在")
	ErrorUploadFileFail    = NewError(20030001, "上传文件失败")
)
