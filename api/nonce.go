package api

import (
	"github.com/xxcheng123/acme-go/constants"
	"github.com/xxcheng123/acme-go/errs"
	"github.com/xxcheng123/acme-go/internal/sender"
)

/**
*	@Author: xxcheng
*	@Email developer@xxcheng.cn
*	@Date: 2024/1/18
 */

// GetNonce Getting a Nonce.
// Nonce is a random string that is used to prevent replay attacks.
// https://datatracker.ietf.org/doc/html/rfc8555#section-7.2
func GetNonce(sender *sender.Sender, nonceURL string) (constants.Nonce, error) {
	resp, err := sender.Head(nonceURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	nonce := resp.Header.Get("Replay-Nonce")
	if nonce == "" {
		return "", errs.GetNonceFail
	}
	return nonce, nil
}
