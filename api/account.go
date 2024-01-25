package api

import (
	"bytes"
	"encoding/json"
	"github.com/xxcheng123/acme-go/constants"
	"github.com/xxcheng123/acme-go/constants/status"
	"github.com/xxcheng123/acme-go/errs"
	"github.com/xxcheng123/acme-go/internal/jws"
	"github.com/xxcheng123/acme-go/internal/sender"
	"io"
	"net/http"
)

/**
*	@Author: xxcheng
*	@Email developer@xxcheng.cn
*	@Date: 2024/1/19
 */

func NewAccount(sender *sender.Sender, NewAccountURL string, account *constants.Account, manager *jws.Manager) (*constants.Account, error) {
	s, _ := json.Marshal(account)
	signed, err := manager.Sign(NewAccountURL, s)
	if err != nil {
		panic(err)
	}
	ss := signed.FullSerialize()
	signedBody := bytes.NewBufferString(ss)
	resp, err := sender.PostJOSE(NewAccountURL, signedBody)
	if resp.StatusCode != http.StatusCreated {
		return nil, errs.StatusNotMatched
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var a constants.Account
	if err = json.Unmarshal(bs, &a); err != nil {
		return nil, err
	}
	if a.Status != status.Valid {
		return nil, errs.CreateNewAccountFailed
	}
	manager.SetKid(resp.Header.Get("Location"))
	nonce := resp.Header.Get("Replay-Nonce")
	if nonce != "" {
		manager.Nonceer.Push(nonce)
	}
	return &a, nil
}
