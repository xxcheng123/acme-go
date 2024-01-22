package constants

/**
*	@Author: xxcheng
*	@Email developer@xxcheng.cn
*	@Date: 2024/1/19
 */

// Order
// An ACME order object represents a client's request for a certificate
// and is used to track the progress of that order through to issuance.
// Example data
//
//	{
//		 "status": "valid",
//		 "expires": "2016-01-20T14:09:07.99Z",
//		 "identifiers": [
//		   { "type": "dns", "value": "www.example.org" },
//		   { "type": "dns", "value": "example.org" }
//		 ],
//		 "notBefore": "2016-01-01T00:00:00Z",
//		 "notAfter": "2016-01-08T00:00:00Z",
//		 "authorizations": [
//		   "https://example.com/acme/authz/PAniVnsZcis",
//		   "https://example.com/acme/authz/r4HqLzrSrpI"
//		 ],
//		 "finalize": "https://example.com/acme/order/TOlocE8rfgo/finalize",
//		 "certificate": "https://example.com/acme/cert/mAt3xBGaobw"
//		}
//
// - status（必需，字符串）：此订单的状态。可能的值包括 "pending"、"ready"、"processing"、"valid" 和 "invalid"。详见 [第7.1.6节](#7.1.6. 状态改变)。
// - expires（可选，字符串）：在服务器将考虑此订单无效之后的时间戳，编码格式遵循 [[RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)] 中指定的格式。此字段对于状态字段中包含 "pending" 或 "valid" 的对象是必需的。
// - identifiers（必需，对象数组）：订单涉及的标识符对象数组。
//   - type（必需，字符串）：标识符的类型。此文档定义了 "dns" 标识符类型。详见 [第9.7.7 节](#9.7.7. 标识符类型) 的注册表，了解其他标识符类型。
//   - value（必需，字符串）：标识符本身。
//
// - notBefore（可选，字符串）：证书中 "notBefore" 字段的请求值，使用 [[RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)] 中定义的日期格式。
// - notAfter（可选，字符串）：证书中 "notAfter" 字段的请求值，使用 [[RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)] 中定义的日期格式。
// - error（可选，对象）：处理订单时发生的错误，如果有的话。此字段结构化为问题文档 [[RFC7807](https://datatracker.ietf.org/doc/html/rfc7807)]。
// - authorizations（必需，字符串数组）：对于待处理的订单，客户端需要在请求的证书发放之前完成的授权（详见 [第7.5节](#7.5. 标识符的授权)），包括客户端过去为订单中指定的标识符完成的未过期授权。所需的授权由服务器策略决定；订单标识符和所需授权之间可能不存在一对一的关系。对于最终订单（处于 "valid" 或 "invalid" 状态），完成的授权。每个条目都是一个URL，可通过POST-as-GET请求从中获取授权。
// - finalize（必需，字符串）：一旦满足订单的所有授权，必须将CSR POST到此URL以完成订单。成功完成的结果将是订单的证书URL的填充。
// - certificate（可选，字符串）：对此订单发放的证书的URL。
type Order struct {
	Status         string        `json:"status"`
	Expires        string        `json:"expires,omitempty"`
	Identifiers    []*Identifier `json:"identifiers"`
	NotBefore      string        `json:"notBefore,omitempty"`
	NotAfter       string        `json:"notAfter,omitempty"`
	Error          any           `json:"error,omitempty"`
	Authorizations []string      `json:"authorizations"`
	Finalize       string        `json:"finalize"`
	Certificate    string        `json:"certificate"`
}
