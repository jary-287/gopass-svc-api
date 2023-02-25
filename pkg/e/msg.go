package e

var MsgFlags = map[int]string{
	ERROR:                   "服务器内部错误",
	ERROR_EXIST_SEVICE:      "服务已经存在",
	SUCCESS:                 "ok",
	ERROR_NOT_EXIST_SERVICE: "不存在这样的服务",
	INVALID_PARAMS:          "参数错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
