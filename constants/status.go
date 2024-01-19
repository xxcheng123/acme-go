package constants

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/19 17:37
 */

// Status List
// https://datatracker.ietf.org/doc/html/rfc8555#section-7.1.6
//
//				pending
//	             |
//	             | Receive
//	             | response
//	             V
//	         processing <-+
//	             |   |    | Server retry or
//	             |   |    | client retry request
//	             |   +----+
//	             |
//	             |
//	 Successful  |   Failed
//	 validation  |   validation
//	   +---------+---------+
//	   |                   |
//	   V                   V
//	 valid              invalid
//
//	 State Transitions for Challenge Objects
const (
	StatusPending    = "pending"
	StatusProcessing = "processing"
)
