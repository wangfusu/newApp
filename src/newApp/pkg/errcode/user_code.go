package errcode

var (
	ErrorGetUserListFail = NewError(20010001, "获取用户列表失败")
	ErrorCreateUserFail  = NewError(20010002, "创建用户失败")
	ErrorUpdateUserFail  = NewError(20010003, "更新用户失败")
	ErrorDeleteUserFail  = NewError(20010004, "删除用户失败")
	ErrorCountUserFail   = NewError(20010005, "统计用户失败")
	ErrorUploadFileFail  = NewError(20030001, "上传文件失败")
)
