package constants

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/19 17:34
 */

// Account
// https://datatracker.ietf.org/doc/html/rfc8555#section-7.1.2
//
//	{
//	   "status": "valid",
//	   "contact": [
//	       "mailto:cert-admin@example.org",
//	       "mailto:admin@example.org"
//	   ],
//	   "termsOfServiceAgreed": true,
//	   "orders": "https://example.com/acme/orders/rzGoeA"
//	}
//
// - status（必需，字符串）：此账户的状态。可能的值有 "valid"、"deactivated" 和 "revoked"。"deactivated" 的值应该用于表示客户端启动的停用，而 "revoked" 应该用于表示服务器启动的停用。参见 [第7.1.6节](#7.1.6. 状态改变)。
// - contact（可选，字符串数组）：服务器可以用来与客户端联系，处理与此账户相关的问题的一组URL。例如，服务器可能希望通知客户端有关服务器启动的吊销或证书过期的情况。有关支持的URL方案的信息，请参见 [第7.3节](#7.3. 账号管理)。
// - termsOfServiceAgreed（可选，布尔值）：在newAccount请求中包含此字段，其值为true，表示客户端同意服务条款。此字段不能由客户端更新。
// - externalAccountBinding（可选，对象）：在newAccount请求中包含此字段表示现有非ACME账户的持有者同意将该账户绑定到此ACME账户。此字段不可由客户端更新（参见 [第7.3.4节](#7.3.4. 外部帐号绑定)）。
// - orders（必需，字符串）：此账户提交的订单列表可以通过POST-as-GET请求从此URL获取，详见 [第7.1.2.1节](#7.1.2.1. 订单列表)。
type Account struct {
	Status                 string                 `json:"status"`
	Contact                []string               `json:"contact,omitempty"`
	TermsOfServiceAgreed   bool                   `json:"termsOfServiceAgreed,omitempty"`
	ExternalAccountBinding ExternalAccountBinding `json:"externalAccountBinding,omitempty"`
	// Orders
	// https://datatracker.ietf.org/doc/html/rfc8555#section-7.1.2.1
	Orders string `json:"orders"`
	// For internal use
	Kid string `json:"-"`
}
