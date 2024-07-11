package api

func SuccessResNone() *ApiRes {
	return &ApiRes{
		Code: 1,
		Msg:  "success",
		Data: nil,
	}
}

func SuccessRes(data any) *ApiRes {
	return &ApiRes{
		Code: 1,
		Msg:  "success",
		Data: data,
	}
}

func ErrorRes(msg string) *ApiRes {
	return &ApiRes{
		Code: 0,
		Msg:  "error",
		Data: nil,
	}
}

type ApiRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
