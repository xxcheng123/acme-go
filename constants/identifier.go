package constants

/**
*	@Author: xxcheng
*	@Email developer@xxcheng.cn
*	@Date: 2024/1/19
 */

// Identifier
// https://datatracker.ietf.org/doc/html/rfc8555#section-9.7.7
// Support List
// |   Label    | Identifier Type | ACME | Reference |
// | :--------: | --------------: | :--: | :-------: |
// |  http-01   |             dns |  Y   | RFC 8555  |
// |   dns-01   |             dns |  Y   | RFC 8555  |
// | tls-sni-01 |        RESERVED |  N   | RFC 8555  |
// | tls-sni-02 |        RESERVED |  N   | RFC 8555  |
type Identifier struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
