package constants

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/18 16:55
 */

// Directory 所有请求 URL 的列表清单以及一个 Meta 信息
// https://datatracker.ietf.org/doc/html/rfc8555#section-7.1.1
// example:
//
//	{
//	   "newNonce": "https://example.com/acme/new-nonce",
//	   "newAccount": "https://example.com/acme/new-account",
//	   "newOrder": "https://example.com/acme/new-order",
//	   "newAuthz": "https://example.com/acme/new-authz",
//	   "revokeCert": "https://example.com/acme/revoke-cert",
//	   "keyChange": "https://example.com/acme/key-change",
//	   "meta": {
//	       "termsOfService": "https://example.com/acme/terms/2017-5-30",
//	       "website": "https://www.example.com/",
//	       "caaIdentities": ["example.com"],
//	       "externalAccountRequired": false
//	   }
//	}
type Directory struct {
	NewNonce   string `json:"newNonce"`
	NewAccount string `json:"newAccount"`
	NewOrder   string `json:"newOrder"`
	// NewAuthz
	// If the ACME server does not implement pre-authorization
	// ([Section 7.4.1](https://datatracker.ietf.org/doc/html/rfc8555#section-7.4.1)),
	// it MUST omit the "newAuthz" field of the directory.
	NewAuthz    string `json:"newAuthz,omitempty"`
	RevokeCert  string `json:"revokeCert"`
	KeyChange   string `json:"keyChange"`
	Meta        Meta   `json:"meta"`
	RenewalInfo string `json:"renewalInfo"`
}

// Meta 元信息
// all of which are OPTIONAL
type Meta struct {
	TermsOfService string   `json:"termsOfService,omitempty"`
	Website        string   `json:"website,omitempty"`
	CaaIdentities  []string `json:"caaIdentities,omitempty"`
	//  ExternalAccountRequired 如果此字段存在且设置为 "true"，则CA要求所有 newAccount 请求都包含一个 "externalAccountBinding" 字段，将新账户与外部账户关联。
	ExternalAccountRequired bool `json:"externalAccountRequired,omitempty"`
}

type Nonce = string
