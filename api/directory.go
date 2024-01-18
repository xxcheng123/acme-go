package api

import (
	"encoding/json"
	"github.com/xxcheng123/acme-go/constants"
	"github.com/xxcheng123/acme-go/internal/sender"
	"io"
)

/**
*	@Author: xxcheng
*	@Email developer@xxcheng.cn
*	@Date: 2024/1/18
 */

func GetDirectory(sender *sender.Sender, directoryURL string) (*constants.Directory, error) {
	resp, err := sender.Get(directoryURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := &constants.Directory{}
	if err := json.Unmarshal(bs, data); err != nil {
		return nil, err
	}
	return data, nil
}
