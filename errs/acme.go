package errs

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/19 16:21
 */

type AcmeError string

const AcmeErrorPrefix = "urn:ietf:params:acme:error:"

const (
	AccountDoesNotExist     AcmeError = "accountDoesNotExist The request specified an account that does not exist"
	AlreadyRevoked          AcmeError = "alreadyRevoked The request specified a certificate to be revoked that has already been revoked"
	BadCSR                  AcmeError = "badCSR The CSR is unacceptable (e.g., due to a short key)"
	BadNonce                AcmeError = "badNonce The client sent an unacceptable anti-replay nonce"
	BadPublicKey            AcmeError = "badPublicKey The JWS was signed by a public key the server does not support"
	BadRevocationReason     AcmeError = "badRevocationReason The revocation reason provided is not allowed by the server"
	BadSignatureAlgorithm   AcmeError = "badSignatureAlgorithm The JWS was signed with an algorithm the server does not support"
	Caa                     AcmeError = "caa Certification Authority Authorization (CAA) records forbid the CA from issuing a certificate"
	Compound                AcmeError = "compound Specific error conditions are indicated in the \"subproblems\" array"
	Connection              AcmeError = "connection The server could not connect to validation target"
	Dns                     AcmeError = "dns There was a problem with a DNS query during identifier validation"
	ExternalAccountRequired AcmeError = "externalAccountRequired The request must include a value for the \"externalAccountBinding\" field"
	IncorrectResponse       AcmeError = "incorrectResponse Response received didn't match the challenge's requirements"
	InvalidContact          AcmeError = "invalidContact A contact URL for an account was invalid"
	Malformed               AcmeError = "malformed The request message was malformed"
	OrderNotReady           AcmeError = "orderNotReady The request attempted to finalize an order that is not ready to be finalized"
	RateLimited             AcmeError = "rateLimited The request exceeds a rate limit"
	RejectedIdentifier      AcmeError = "rejectedIdentifier The server will not issue certificates for the identifier"
	ServerInternal          AcmeError = "serverInternal The server experienced an internal error"
	Tls                     AcmeError = "tls The server received a TLS error during validation"
	Unauthorized            AcmeError = "unauthorized The client lacks sufficient authorization"
	UnsupportedContact      AcmeError = "unsupportedContact A contact URL for an account used an unsupported protocol scheme"
	UnsupportedIdentifier   AcmeError = "unsupportedIdentifier An identifier is of an unsupported type"
	UserActionRequired      AcmeError = "userActionRequired Visit the \"instance\" URL and take actions specified there"
)

type Problem struct {
	Type       AcmeError     `json:"type"`
	Detail     string        `json:"detail"`
	Status     int           `json:"status"`
	SubProblem []*SubProblem `json:"subproblem,omitempty"`
}

func (p *Problem) Error() string {
	return p.Detail
}
