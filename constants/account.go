package constants

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/19 17:34
 */

// Account
// https://datatracker.ietf.org/doc/html/rfc8555#section-7.1.2
type Account struct {
	Status                 string   `json:"status"`
	Contact                []string `json:"contact,omitempty"`
	TermsOfServiceAgreed   bool     `json:"termsOfServiceAgreed,omitempty"`
	ExternalAccountBinding any      `json:"externalAccountBinding,omitempty"`
	Orders                 string   `json:"orders"`
}
