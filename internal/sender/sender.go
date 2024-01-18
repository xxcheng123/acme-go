package sender

import (
	"fmt"
	acme_go "github.com/xxcheng123/acme-go"
	"github.com/xxcheng123/acme-go/errs"
	"io"
	"net/http"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/18 17:10
 */
const defaultUserAgentPrefix = "acme-go"

func defaultUserAgentFormat() string {
	return fmt.Sprintf("%s/%s", defaultUserAgentPrefix, acme_go.Version)
}

type Sender struct {
	client        *http.Client
	userAgentFunc func() string
}
type NewSenderOption func(s *Sender)

var globalSender = NewSender()

func GetSender() *Sender {
	return globalSender
}
func NewSender(opts ...NewSenderOption) *Sender {
	s := &Sender{
		client:        &http.Client{},
		userAgentFunc: defaultUserAgentFormat,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// do handle something after response
func (s *Sender) do(req *http.Request) (*http.Response, error) {
	if req == nil {
		s.client = &http.Client{}
	}
	return s.client.Do(req)
}

// newRequest wrap something before request
func (s *Sender) newRequest(name string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(name, url, body)
	if err != nil {
		return nil, err
	}

	// according to ref8555
	// ACME clients MUST send a User-Agent header field
	// https://datatracker.ietf.org/doc/html/rfc8555#section-6.1
	if s.userAgentFunc != nil {
		req.Header.Set("User-Agent", s.userAgentFunc())
	} else {
		return nil, errs.MissFunction
	}
	return req, err
}

func (s *Sender) Get(url string) (*http.Response, error) {
	req, err := s.newRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return s.do(req)
}
func (s *Sender) Head(url string) (*http.Response, error) {
	req, err := s.newRequest(http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}
	return s.do(req)
}
