package client

import (
	"github.com/xxcheng123/acme-go/api"
	"github.com/xxcheng123/acme-go/constants"
	"github.com/xxcheng123/acme-go/internal/sender"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/18 17:19
 */

type Client struct {
	directoryURL string
	sender       *sender.Sender
	directory    *constants.Directory
}

func NewClient(directoryURL string) (*Client, error) {
	sdr := sender.NewSender()
	directory, err := api.GetDirectory(sdr, directoryURL)
	if err != nil {
		return nil, err
	}
	return &Client{
		directoryURL: directoryURL,
		sender:       sdr,
		directory:    directory,
	}, nil
}
