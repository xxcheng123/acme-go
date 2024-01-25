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
	NotSupportedCrypto
	CreateNewAccountFailed
	StatusNotMatched
	NotAgree
)

const defaultErrorMessage = "unknown error"

var codeMapErrorMessage = map[Code]string{
	MissFunction:           "Miss Function",
	GetNonceFail:           "Get Nonce Fail",
	NotSupportedCrypto:     "Not Supported Crypto",
	CreateNewAccountFailed: "Create New Account Failed",
	StatusNotMatched:       "Status Not Matched",
	NotAgree:               "you must agree to the terms of service",
}

func (c Code) Error() string {
	if msg, ok := codeMapErrorMessage[c]; ok {
		return msg
	}
	return defaultErrorMessage
}
