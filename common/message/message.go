package message

const (
	LoginMesType       = "LoginMes"
	RegisterMesType    = "RegisterMes"
	LoginResMesType    = "LoginResMes"
	RegisterResMesType = "RegisterResMes"

	SuccessCode           = 200
	UserOrPasswordInvalid = 401
	UserNotExist          = 402
	UserRegisterExisted   = 403
	ServerError           = 500
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type RegisterMes struct {
	UserVO *User `json:user`
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:error`
}
