package request

import (
	"github.com/xxcheng123/acme-go/constants"
	"time"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/30 11:15
 */

// Order
// Apply Certificate
// https://datatracker.ietf.org/doc/html/rfc8555#section-7.4
type Order struct {
	Identifiers []constants.Identifier `json:"identifiers"`
	NotBefore   time.Time
	NotAfter    time.Time
}

func NewOrder(identifiers []constants.Identifier, notBefore time.Time, notAfter time.Time) *Order {
	return &Order{
		Identifiers: identifiers,
		NotBefore:   notBefore,
		NotAfter:    notAfter,
	}
}
