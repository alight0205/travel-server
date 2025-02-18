package res

type ErrorCode int

const (
	SettingsError ErrorCode = 1001 // 设置错误
	ArgumentError ErrorCode = 1002 // 参数错误
	DatabaseError ErrorCode = 1003 // 数据库错误
	RedisError    ErrorCode = 1004 // redis 错误
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "设置错误",
		ArgumentError: "参数错误",
		DatabaseError: "数据库错误",
		RedisError:    "redis 错误",
	}
)
