package encryption

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/xxcheng123/acme-go/errs"
	"os"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/26 15:23
 */

type KeyType string

const (
	EC256   KeyType = "EC256"
	EC384   KeyType = "EC384"
	RSA2048 KeyType = "RSA2048"
	RSA3072 KeyType = "RSA3072"
	RSA4096 KeyType = "RSA4096"
	RSA8192 KeyType = "RSA8192"
)

func GeneratePrivateKey(keyType KeyType) (crypto.PrivateKey, error) {
	switch keyType {
	// https://datatracker.ietf.org/doc/html/rfc7518
	case EC256:
		return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case EC384:
		return ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case RSA2048:
		return rsa.GenerateKey(rand.Reader, 2048)
	case RSA3072:
		return rsa.GenerateKey(rand.Reader, 3072)
	case RSA4096:
		return rsa.GenerateKey(rand.Reader, 4096)
	case RSA8192:
		return rsa.GenerateKey(rand.Reader, 8192)
	default:
		return nil, errs.NotSupportedCrypto
	}
}

func SavePrivateKey(key crypto.PrivateKey, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	pemKey, err := convert2PEMBlock(key)
	if err != nil {
		return err
	}
	err = pem.Encode(f, pemKey)
	if err != nil {
		return err
	}
	return nil
}
func ConvertPrivateKey(key crypto.PrivateKey) ([]byte, error) {
	pemKey, err := convert2PEMBlock(key)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(pemKey), nil
}
func ParsePrivateKey(bs []byte) (crypto.PrivateKey, error) {
	keyBlock, _ := pem.Decode(bs)
	switch keyBlock.Type {
	case "RSA PRIVATE KEY":
		return x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	case "EC PRIVATE KEY":
		return x509.ParseECPrivateKey(keyBlock.Bytes)
	}

	return nil, errs.NotSupportedCrypto
}
func OpenPrivateKey(path string) (crypto.PrivateKey, error) {
	keyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParsePrivateKey(keyBytes)
}

func convert2PEMBlock(data any) (*pem.Block, error) {
	var p *pem.Block
	switch key := data.(type) {
	case *ecdsa.PrivateKey:
		keyBytes, _ := x509.MarshalECPrivateKey(key)
		p = &pem.Block{Type: "EC PRIVATE KEY", Bytes: keyBytes}
	case *rsa.PrivateKey:
		p = &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}
	case *x509.CertificateRequest:
		p = &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: key.Raw}
	default:
		return nil, errs.NotSupportedCrypto
	}
	return p, nil
}
