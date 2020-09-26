package utils

// 常量定义规范
// 1、奇数 => error 偶数 => success
const (
	// Server code
	SERVER_OK    = 10000
	SERVER_ERROR = 10001

	// User code
	USER_OK               = 20000
	USER_EXIST            = 20001
	USER_EXIST_USED       = 20003
	USER_REGISTER_SUCCESS = 20004
	USER_REGISTER_FAIL    = 20005
	USER_NO_EXIST         = 20007
)

var codeMessage = map[int]string{
	// Server code
	SERVER_OK:    "ok",
	SERVER_ERROR: "error",

	// User code
	USER_OK:               "可以使用",
	USER_EXIST:            "用户已存在",
	USER_EXIST_USED:       "用户名已被使用",
	USER_REGISTER_SUCCESS: "注册成功",
	USER_REGISTER_FAIL:    "注册失败",
	USER_NO_EXIST:         "用户不存在",
}

func GetMessage(code int) string {
	return codeMessage[code]
}
