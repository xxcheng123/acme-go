package client

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/25 16:49
 */

func TestNewClient(t *testing.T) {
	k, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Error(err)
		return
	}
	client, err := NewClient("https://acme-v02.api.letsencrypt.org/directory", k)
	if err != nil {
		t.Error(err)
		return
	}
	err = client.CreateAccount("hello@example.com", true)
	if err != nil {
		t.Error(err)
	}
}
