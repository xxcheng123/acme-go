package constants

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/18 16:55
 */

// Directory 所有请求 URL 的列表清单以及一个 Meta 信息
// https://datatracker.ietf.org/doc/html/rfc8555#section-7.1.1
type Directory struct {
	NewNonce    string `json:"newNonce"`
	NewAccount  string `json:"newAccount"`
	NewOrder    string `json:"newOrder"`
	NewAuthz    string `json:"newAuthz"`
	RevokeCert  string `json:"revokeCert"`
	KeyChange   string `json:"keyChange"`
	Meta        Meta   `json:"meta"`
	RenewalInfo string `json:"renewalInfo"`
}

// Meta 元信息
type Meta struct {
	TermsOfService string   `json:"termsOfService"`
	Website        string   `json:"website"`
	CaaIdentities  []string `json:"caaIdentities"`
	//  ExternalAccountRequired 如果此字段存在且设置为 "true"，则CA要求所有 newAccount 请求都包含一个 "externalAccountBinding" 字段，将新账户与外部账户关联。
	ExternalAccountRequired bool `json:"externalAccountRequired"`
}

type Nonce = string
