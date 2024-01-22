package constants

/**
*	@Author: xxcheng
*	@Email developer@xxcheng.cn
*	@Date: 2024/1/19
 */

// Identifier
// https://datatracker.ietf.org/doc/html/rfc8555#section-9.7.7

type Identifier struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
