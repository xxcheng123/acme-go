package nonceer

import (
	"github.com/xxcheng123/acme-go/api"
	"github.com/xxcheng123/acme-go/errs"
	"github.com/xxcheng123/acme-go/internal/sender"
	"sync"
	"time"
)

/**
*	@Author: xxcheng
*	@Email developer@xxcheng.cn
*	@Date: 2024/1/18
 */

// defaultCacheSize
const defaultCacheSize = 8
const defaultLoopNonceSleep = time.Second * 1

// Nonceer a manager to apply nonce
type Nonceer struct {
	sender         *sender.Sender
	caches         chan string
	nonceURL       string
	loopNonceSleep time.Duration
	mu             sync.Mutex
}
type Nonce = string
type NewNonceerOption func(nonceer *Nonceer)

func SetCacheSize(size int) NewNonceerOption {
	return func(nonceer *Nonceer) {
		nonceer.caches = make(chan string, size)
	}
}

func NewNonceer(sender *sender.Sender, nonceURL string, opts ...NewNonceerOption) (*Nonceer, error) {
	nonceer := &Nonceer{
		sender:         sender,
		nonceURL:       nonceURL,
		caches:         make(chan string, defaultCacheSize),
		loopNonceSleep: defaultLoopNonceSleep,
	}
	for _, opt := range opts {
		opt(nonceer)
	}
	return nonceer, nil
}
func (n *Nonceer) Get() (Nonce, error) {
	go n.newNonce()
	select {
	case nonce := <-n.caches:
		return nonce, nil
	case <-time.After(time.Second * 10):
		return "", errs.GetNonceFail
	}
}

// newNonce To obtain the nonce multiple times, prefetch some Nonce.
func (n *Nonceer) newNonce() {
	n.mu.Lock()
	defer n.mu.Unlock()
	for len(n.caches) < cap(n.caches) {
		nonce, err := api.GetNonce(n.sender, n.nonceURL)
		if err != nil {
			break
		}
		n.caches <- nonce
		if len(n.caches) == cap(n.caches) {
			break
		}
		time.Sleep(n.loopNonceSleep)
	}
}
