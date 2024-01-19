package errs

import "fmt"

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
	Identifier Identifier `json:"identifier,omitempty"`
}

func (s *SubProblem) Error() string {
	return fmt.Sprintf("error:%s,detail:%s", s.Type, s.Detail)
}

// Identifier
// https://datatracker.ietf.org/doc/html/rfc8555#section-9.7.7
// It can only be present in SubProblems.
// SubProblems need not all have the same type, and they do not need to
// match the top level type.
type Identifier struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
