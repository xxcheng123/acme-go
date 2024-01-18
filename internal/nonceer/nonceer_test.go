package nonceer

import (
	"fmt"
	acme_go "github.com/xxcheng123/acme-go"
	"github.com/xxcheng123/acme-go/internal/sender"
	"testing"
)

/**
*	@Author: xxcheng
*	@Email developer@xxcheng.cn
*	@Date: 2024/1/18
 */

func TestNewNonceer(t *testing.T) {
	nonceer, err := NewNonceer(sender.NewSender(), acme_go.DefaultNewNonceURL)
	if err != nil {
		t.Error(err)
	}
	nonce, err := nonceer.Get()
	if err != nil {
		t.Error(err)
	}
	if nonce == "" {
		t.Error("nonce is empty")
	}
	fmt.Println(nonce)
}
