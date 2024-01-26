package core

import (
	"github.com/xxcheng123/acme-go/constants"
	"github.com/xxcheng123/acme-go/internal/jws"
	"github.com/xxcheng123/acme-go/internal/sender"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/26 13:32
 */

// Core
// For reuse.
type Core struct {
	DirectoryURL string
	Sender       *sender.Sender
	Directory    *constants.Directory
	JWSManager   *jws.Manager
}
type CloneOpt func(*Core) *Core

func (c *Core) Clone(opts ...CloneOpt) *Core {
	nc := &Core{
		DirectoryURL: c.DirectoryURL,
		Sender:       c.Sender,
		Directory:    c.Directory,
		JWSManager:   c.JWSManager,
	}
	for _, opt := range opts {
		nc = opt(nc)
	}
	return nc
}
