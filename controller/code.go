package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeServerBusy
	CodeError

	CodeInvalidToken
	CodeNeedLogin

	CodeInvalidFile
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:      "success",
	CodeInvalidParam: "invalid param",
	CodeServerBusy:   "server busy",
	CodeError:        "error",

	CodeInvalidToken: "invalid token",
	CodeNeedLogin:    "need login",

	CodeInvalidFile: "invalid file",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeError]
	}
	return msg
}
