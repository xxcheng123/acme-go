package client

import (
	"crypto"
	"encoding/json"
	"fmt"
	"github.com/xxcheng123/acme-go/account"
	"github.com/xxcheng123/acme-go/api"
	"github.com/xxcheng123/acme-go/constants/request"
	"github.com/xxcheng123/acme-go/core"
	"github.com/xxcheng123/acme-go/errs"
	"github.com/xxcheng123/acme-go/internal/encryption"
	"github.com/xxcheng123/acme-go/internal/jws"
	"github.com/xxcheng123/acme-go/internal/nonceer"
	"github.com/xxcheng123/acme-go/internal/sender"
	"os"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/18 17:19
 */

type Client struct {
	Core *core.Core
}
type PerClient struct {
	DirectoryURL string `json:"directoryURL"`
	PrivateKey   string `json:"privateKey"`
}

func NewClient(directoryURL string, privateKey crypto.PrivateKey) (*Client, error) {
	sdr := sender.NewSender()
	directory, err := api.GetDirectory(sdr, directoryURL)

	if err != nil {
		return nil, err
	}
	newNonceer, err := nonceer.NewNonceer(sdr, directory.NewNonce)
	if err != nil {
		return nil, err
	}
	jwtManager, err := jws.NewManager(privateKey, newNonceer, "")
	if err != nil {
		return nil, err
	}
	return &Client{
		Core: &core.Core{
			DirectoryURL: directoryURL,
			Sender:       sdr,
			Directory:    directory,
			JWSManager:   jwtManager,
		},
	}, nil
}

func (c *Client) CreateAccount(email string, agree bool) (*account.Account, error) {
	if !agree {
		return nil, errs.NotAgree
	}
	accReq := &request.Account{
		Contact: []string{
			fmt.Sprintf("mailto:%s", email),
		},
		TermsOfServiceAgreed: agree,
	}
	acc, err := api.NewAccount(c.Core.Sender, c.Core.JWSManager, c.Core.Directory.NewAccount, accReq)
	if err != nil {
		return nil, err
	}
	nc := c.Core.Clone()
	nm := nc.JWSManager.Clone(jws.CustomKid(acc.Kid))
	nc.JWSManager = nm
	accService := account.NewAccount(nc, acc)
	return accService, nil
}

func (c *Client) Persist(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	bs, err := encryption.ConvertPrivateKey(c.Core.JWSManager.GetPrivateKey())
	if err != nil {
		return err
	}
	p := &PerClient{
		DirectoryURL: c.Core.DirectoryURL,
		PrivateKey:   string(bs),
	}
	if bs, err = json.Marshal(p); err != nil {
		return err
	}
	err = os.WriteFile(path, bs, 0222)
	if err != nil {
		return err
	}
	return nil
}
