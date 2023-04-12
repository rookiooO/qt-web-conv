package server

func Exception(code int, message string) {
	panic(JsonResult{
		Code: code,
		Msg:  message,
	})
}

func ExceptionWithCode(code int) {
	if msg, ok := StatusMsg[code]; ok {
		panic(JsonResult{
			Code: code,
			Msg:  msg,
		})
	} else {
		panic(JsonResult{
			Code: code,
			Msg:  "未知错误",
		})
	}
}

func ExceptionWithMsg(message string) {
	panic(JsonResult{
		Code: 0,
		Msg:  message,
	})
}
