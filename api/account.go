package api

import (
	"encoding/json"
	"github.com/xxcheng123/acme-go/constants"
	"github.com/xxcheng123/acme-go/constants/request"
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

func NewAccount(sender *sender.Sender, manager *jws.Manager, NewAccountURL string, account *request.Account) (*constants.Account, error) {
	s, _ := json.Marshal(account)
	resp, err := postJose(sender, manager, NewAccountURL, s)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		var e errs.Problem
		if err = json.Unmarshal(bs, &e); err == nil {
			return nil, &e
		}
		return nil, errs.StatusNotMatched
	}
	var a constants.Account
	if err = json.Unmarshal(bs, &a); err != nil {
		return nil, err
	}
	if a.Status != status.Valid {
		return nil, errs.CreateNewAccountFailed
	}
	if kid := resp.Header.Get("Location"); kid != "" {
		a.Kid = kid
	}
	return &a, nil
}
