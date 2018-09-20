package rpcModels

// rpc登录参数
type LoginParam struct {
	Username	string		`json:"USERNAME"`
	Password	string		`json:"PASSWORD"`
}

// rpc登录返回信息
type LoginReply struct {
	Id			int
	Username	string
	Nickname	string
	baseReply
}
