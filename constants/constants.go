package constants

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/18 16:55
 */

// Nonce
// https://datatracker.ietf.org/doc/html/rfc8555#autoid-13
// 服务器可以用它来检测未经授权的重放攻击在未来的客户端请求中。
// 服务器必须以这样的方式生成 Replay-Nonce 头字段中提供的值，
// 以使它们对于每个消息具有高概率的唯一性，并且对于服务器以外的任何人都是不可预测的。
type Nonce = string
