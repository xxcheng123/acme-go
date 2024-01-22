package errs

import (
	"fmt"
	"github.com/xxcheng123/acme-go/constants"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/19 16:28
 */

// SubProblem
// https://datatracker.ietf.org/doc/html/rfc8555#section-6.7.1
// example json format:
//
//	 {
//	    "type": "urn:ietf:params:acme:error:malformed",
//	    "detail": "Invalid underscore in DNS name \"_example.org\"",
//	    "identifier": {
//	        "type": "dns",
//	        "value": "_example.org"
//	    }
//	}
type SubProblem struct {
	// Type prefix[urn:ietf:params:acme:error:]
	Type AcmeError `json:"type,omitempty"`
	// Detail
	Detail string `json:"detail,omitempty"`
	// Identifier
	// ACME clients may choose to use the "identifier" field of a subProblem
	// as a hint that an operation would succeed if that identifier were
	// omitted.
	Identifier constants.Identifier `json:"identifier,omitempty"`
}

func (s *SubProblem) Error() string {
	return fmt.Sprintf("error:%s,detail:%s", s.Type, s.Detail)
}
