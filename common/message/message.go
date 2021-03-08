package message

const (
	LoginMesType            = "LoginMes"
	RegisterMesType         = "RegisterMes"
	LoginResMesType         = "LoginResMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMesType"
	SmsMesType              = "SmsMesType"
	SmsTransferMesType      = "SmsTransferMesType"

	SmsMesPrivateType         = "SmsMesPrivateType"
	SmsTransferPrivateMesType = "SmsTransferPrivateMesType"

	UserOffline = 1
	UserOnline  = 2
	UserBusy    = 3

	SuccessCode           = 200
	UserOrPasswordInvalid = 401
	UserNotExist          = 402
	UserRegisterExisted   = 403
	UserNotOnline         = 404
	ServerError           = 500
)

type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:status`
}

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
	Code    int    `json:"code"`
	Error   string `json:"error"`
	UserIds []int  `json:"userIds"`
}

type RegisterMes struct {
	UserVO *User `json:user`
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:error`
}

type SmsMes struct {
	User
	Content string
}

type SmsTransferMes struct {
	User
	Content string
	Code    int    `json:"code"`
	Error   string `json:error`
}
