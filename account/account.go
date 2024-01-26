package account

import (
	"github.com/xxcheng123/acme-go/constants"
	"github.com/xxcheng123/acme-go/core"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/26 12:00
 */

type Account struct {
	Acc  *constants.Account
	Core *core.Core
}

func NewAccount(core *core.Core, acc *constants.Account) *Account {
	return &Account{Core: core, Acc: acc}
}
