package api

import (
	"github.com/stretchr/testify/assert"
	acme_go "github.com/xxcheng123/acme-go"
	"github.com/xxcheng123/acme-go/internal/sender"
	"testing"
)

/**
*	@Author: xxcheng
*	@Email developer@xxcheng.cn
*	@Date: 2024/1/18
 */
var sdr *sender.Sender

func init() {
	sdr = sender.NewSender()
}

func Test_GetDirectory(t *testing.T) {
	_, err := GetDirectory(sdr, acme_go.DefaultAcmeDirectoryURL)
	assert.Equal(t, nil, err)
}
func Test_GetNonce(t *testing.T) {
	_, err := GetNonce(sdr, acme_go.DefaultNewNonceURL)
	assert.Equal(t, nil, err)
}
