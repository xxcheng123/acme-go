package errs

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/18 17:25
 */

type Code uint32

const (
	MissFunction Code = iota + 1000
	GetNonceFail
)

const defaultErrorMessage = "unknown error"

var codeMapErrorMessage = map[Code]string{
	MissFunction: "Miss Function",
	GetNonceFail: "Get Nonce Fail",
}

func (c Code) Error() string {
	if msg, ok := codeMapErrorMessage[c]; ok {
		return msg
	}
	return defaultErrorMessage
}
