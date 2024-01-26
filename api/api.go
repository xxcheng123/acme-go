package api

import (
	"bytes"
	"github.com/xxcheng123/acme-go/internal/jws"
	"github.com/xxcheng123/acme-go/internal/sender"
	"net/http"
)

/**
*	@Author: xxcheng
*	@Email developer@xxcheng.cn
*	@Date: 2024/1/19
 */

func postJose(sender *sender.Sender, manager *jws.Manager, url string, content []byte) (*http.Response, error) {
	signed, err := manager.Sign(url, content)
	if err != nil {
		panic(err)
	}
	ss := signed.FullSerialize()
	signedBody := bytes.NewBufferString(ss)
	resp, err := sender.PostJOSE(url, signedBody)
	if err != nil {
		return nil, err
	}
	nonce := resp.Header.Get("Replay-Nonce")
	if nonce != "" {
		manager.Nonceer.Push(nonce)
	}
	//if kid := resp.Header.Get("Location"); kid != "" {
	//	manager.SetKid(kid)
	//}
	return resp, err
}
