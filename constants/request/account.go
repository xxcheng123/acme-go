package request

import "github.com/xxcheng123/acme-go/constants"

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/25 10:37
 */

// Account for request struct
type Account struct {
	Contact                []string                         `json:"contact,omitempty"`
	TermsOfServiceAgreed   bool                             `json:"termsOfServiceAgreed,omitempty"`
	OnlyReturnExisting     bool                             `json:"onlyReturnExisting,omitempty"`
	ExternalAccountBinding constants.ExternalAccountBinding `json:"externalAccountBinding,omitempty"`
}
