package client

import (
	"fmt"
	"github.com/xxcheng123/acme-go/api"
	"github.com/xxcheng123/acme-go/constants"
	"github.com/xxcheng123/acme-go/errs"
)

/**
*	@Author: xxcheng
*	@Email developer@xxcheng.cn
*	@Date: 2024/1/19
 */

func (c *Client) CreateAccount(email string, agree bool) error {
	if !agree {
		return errs.NotAgree
	}
	account := &constants.Account{
		Contact: []string{
			fmt.Sprintf("mailto:%s", email),
		},
		TermsOfServiceAgreed: agree,
	}
	acc, err := api.NewAccount(c.Sender, c.Directory.NewAccount, account, c.JWSManager)
	if err != nil {
		return err
	}
	c.Account = acc
	return nil
}
