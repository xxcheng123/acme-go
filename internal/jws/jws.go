package jws

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"github.com/go-jose/go-jose/v3"
	"github.com/xxcheng123/acme-go/constants"
	"github.com/xxcheng123/acme-go/errs"
	"github.com/xxcheng123/acme-go/internal/nonceer"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/25 11:15
 */

// JWS
// https://datatracker.ietf.org/doc/html/rfc8555#section-6.2
type JWS struct {
	Alg   string          `json:"alg"`
	Nonce constants.Nonce `json:"nonce"`
	// Url
	// https://datatracker.ietf.org/doc/html/rfc8555#section-6.4.1
	// The "url" header parameter MUST be carried
	// in the protected header of the JWS.
	Url string `json:"url"`
	// The "jwk" and "kid" fields are mutually exclusive.
	Jwk string `json:"jwk,omitempty"`
	Kid string `json:"kid,omitempty"`
}

type Manager struct {
	Nonceer    *nonceer.Nonceer
	Alg        jose.SignatureAlgorithm
	privateKey crypto.PrivateKey
	kid        string
}

func (m *Manager) SetKid(kid string) {
	m.kid = kid
}
func NewManager(privateKey crypto.PrivateKey, nonceer *nonceer.Nonceer, kid string) (*Manager, error) {
	var alg jose.SignatureAlgorithm
	switch k := privateKey.(type) {
	case *rsa.PrivateKey:
		alg = jose.RS256
	case *ecdsa.PrivateKey:
		if k.Curve == elliptic.P256() {
			alg = jose.ES256
		} else if k.Curve == elliptic.P384() {
			alg = jose.ES384
		}
	default:
		return nil, errs.NotSupportedCrypto
	}
	return &Manager{
		privateKey: privateKey,
		Nonceer:    nonceer,
		Alg:        alg,
		kid:        kid,
	}, nil
}

// Sign
// Generate a JWS object
// Document see https://pkg.go.dev/github.com/go-jose/go-jose/v3
func (m *Manager) Sign(url string, payload []byte) (*jose.JSONWebSignature, error) {
	signKey := jose.SigningKey{
		Algorithm: m.Alg,
		Key:       jose.JSONWebKey{Key: m.privateKey, KeyID: m.kid},
	}
	signOptions := jose.SignerOptions{
		NonceSource: m.Nonceer,
		ExtraHeaders: map[jose.HeaderKey]interface{}{
			"url": url,
		},
	}
	if m.kid == "" {
		signOptions.EmbedJWK = true
	}
	signer, err := jose.NewSigner(signKey, &signOptions)
	if err != nil {
		return nil, err
	}
	return signer.Sign(payload)
}
