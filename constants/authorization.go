package constants

/**
*	@Author: xxcheng
*	@Email developer@xxcheng.cn
*	@Date: 2024/1/19
 */

// Authorization
// https://datatracker.ietf.org/doc/html/rfc8555#section-7.1.4
// Example data:
//
//	{
//	 "status": "valid",
//	 "expires": "2015-03-01T14:09:07.99Z",
//
//	 "identifier": {
//	   "type": "dns",
//	   "value": "www.example.org"
//	 },
//
//	 "challenges": [
//	   {
//	     "url": "https://example.com/acme/chall/prV_B7yEyA4",
//	     "type": "http-01",
//	     "status": "valid",
//	     "token": "DGyRejmCefe7v4NfDGDKfA",
//	     "validated": "2014-12-01T12:05:58.16Z"
//	   }
//	 ],
//
//	 "wildcard": false
//	}
//
// - identifier（必需，对象）：帐户被授权代表的标识符。
//   - type（必需，字符串）：标识符的类型（参见下文和 [第9.7.7节](#9.7.7. 标识符类型)）。
//   - value（必需，字符串）：标识符本身。
//
// - status（必需，字符串）：此授权的状态。可能的值包括 "pending"、"valid"、"invalid"、"deactivated"、"expired" 和 "revoked"。详见 [第7.1.6节](#7.1.6. 状态改变)。
// - expires（可选，字符串）：服务器将在此授权无效之后考虑的时间戳，编码格式遵循 [[RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)] 中指定的格式。对于状态字段中包含 "valid" 的对象，此字段是必需的。
// - challenges（必需，对象数组）：对于待处理的授权，客户端可以履行以证明对标识符的所有权的挑战。对于有效的授权，已验证的挑战。对于无效的授权，尝试并失败的挑战。每个数组条目都是一个对象，其中包含验证挑战所需的参数。客户端应尝试履行这些挑战中的一个，服务器应考虑这些挑战中的任何一个足以使授权有效。
// - wildcard（可选，布尔值）：对于作为 newOrder 请求的结果创建的授权，该请求包含一个值为通配符域名的 DNS 标识符，此字段必须存在且为 true。对于其他授权，它必须不存在。通配符域名在 [第7.1.3节](#7.1.3. 订单对象) 中有描述。
type Authorization struct {
	Identifier `json:"identifier"`
	Status     string       `json:"status"`
	Expires    string       `json:"expires,omitempty"`
	Challenges []*Challenge `json:"challenges"`
	Wildcard   bool         `json:"wildcard,omitempty"`
}
