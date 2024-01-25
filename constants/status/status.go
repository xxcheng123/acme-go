package status

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/25 17:00
 */

// Status List
// https://datatracker.ietf.org/doc/html/rfc8555#section-7.1.6
//
//					pending
//		             |
//		             | Receive
//		             | response
//		             V
//		         processing <-+
//		             |   |    | Server retry or
//		             |   |    | client retry request
//		             |   +----+
//		             |
//		             |
//		 Successful  |   Failed
//		 validation  |   validation
//		   +---------+---------+
//		   |                   |
//		   V                   V
//		 valid              invalid
//
//		 State Transitions for Challenge Objects
//
//
//	                      pending --------------------+
//	                        |                        |
//	      Challenge failure |                        |
//	             or         |                        |
//	            Error       |  Challenge valid       |
//	              +---------+---------+              |
//	              |                   |              |
//	              V                   V              |
//	           invalid              valid            |
//	                                  |              |
//	                                  |              |
//	                                  |              |
//	                   +--------------+--------------+
//	                   |              |              |
//	                   |              |              |
//	            Server |       Client |   Time after |
//	            revoke |   deactivate |    "expires" |
//	                   V              V              V
//	                revoked      deactivated      expired
//
//	               State Transitions for Authorization Objects
const (
	Pending     = "pending"
	Ready       = "ready"
	Processing  = "processing"
	Valid       = "valid"
	Invalid     = "invalid"
	Expired     = "expired"
	Deactivated = "deactivated"
	Revoked     = "revoked"
)
