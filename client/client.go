package client

import (
	"crypto"
	"github.com/xxcheng123/acme-go/api"
	"github.com/xxcheng123/acme-go/constants"
	"github.com/xxcheng123/acme-go/internal/jws"
	"github.com/xxcheng123/acme-go/internal/nonceer"
	"github.com/xxcheng123/acme-go/internal/sender"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/18 17:19
 */

type Client struct {
	DirectoryURL string
	Sender       *sender.Sender
	Directory    *constants.Directory
	Nonceer      *nonceer.Nonceer
	JWSManager   *jws.Manager
	Account      *constants.Account
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
		DirectoryURL: directoryURL,
		Sender:       sdr,
		Directory:    directory,
		Nonceer:      newNonceer,
		JWSManager:   jwtManager,
		Account:      nil,
	}, nil
}
