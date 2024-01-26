package nonceer

import (
	"github.com/xxcheng123/acme-go/constants"
	"github.com/xxcheng123/acme-go/errs"
	"github.com/xxcheng123/acme-go/internal/sender"
	"sync"
)

/**
*	@Author: xxcheng
*	@Email developer@xxcheng.cn
*	@Date: 2024/1/18
 */

// defaultSize
const defaultSize = 8

// Nonceer a manager to apply nonce
type Nonceer struct {
	sender   *sender.Sender
	list     []string
	nonceURL string
	mu       sync.Mutex
}
type NewNonceerOption func(nonceer *Nonceer)

func SetDefaultSize(size int) NewNonceerOption {
	return func(nonceer *Nonceer) {
		nonceer.list = make([]string, 0, defaultSize)
	}
}

func NewNonceer(sender *sender.Sender, nonceURL string, opts ...NewNonceerOption) (*Nonceer, error) {
	nonceer := &Nonceer{
		sender:   sender,
		nonceURL: nonceURL,
		list:     make([]string, 0, defaultSize),
	}
	for _, opt := range opts {
		opt(nonceer)
	}
	return nonceer, nil
}
func (n *Nonceer) Get() (constants.Nonce, error) {
	if nonce, ok := n.Pop(); ok {
		return nonce, nil
	}
	return GetNonce(n.sender, n.nonceURL)
}
func (n *Nonceer) Pop() (string, bool) {
	n.mu.Lock()
	defer n.mu.Unlock()
	if len(n.list) > 0 {
		nonce := n.list[0]
		n.list = n.list[1:]
		return nonce, true
	}
	return "", false
}
func (n *Nonceer) Push(nonce string) bool {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.list = append(n.list, nonce)
	return true
}

// Nonce Implement jose.NonceSource.
func (n *Nonceer) Nonce() (string, error) {
	return n.Get()
}

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
