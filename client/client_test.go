package client

import (
	"fmt"
	"github.com/xxcheng123/acme-go/internal/encryption"
	"testing"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/25 16:49
 */
const c = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAv+8Z2sCYhXEC0+IiVbtS\nFlZ51ci7aqfTw/adMkVDvWJvKm5OjB8Ej/Z6ITCjyxZGe/LTOmJ1s5l3qOesqKPP\ngTzoinc03SCEn0FdgKj8tv/2UYHcntks9SbrZeuleSatB5tDr+r1PCwFyDqUgUoc\nmCzzb9me6FDhtfRuMNcqCJcf6mmji83OYz+dc/PpGA7urEa5ED4IU06UZ0LtUyeh\nHezfXP5K5p84jQpEEvPKo1JqfZGCP4LRMRY7vXurR6yIoyrydO2T6w5SjK7b/3yq\nDN2bwFujpPryYhyjpn5Mcp0p9YENwAh05fC++/AoUKYq6EPrZ1fIMWJ1pw2eZBHP\nVQIDAQAB\n-----END PUBLIC KEY-----\n"

func TestNewClient(t *testing.T) {
	//k, err := encryption.GeneratePrivateKey(encryption.EC256)
	k, err := encryption.ParsePrivateKey([]byte(c))
	//_ = encryption.SavePrivateKey(k, "_.key")
	if err != nil {
		t.Error(err)
		return
	}
	//k, err := rsa.GenerateKey(rand.Reader, 2048)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	client, err := NewClient("https://acme-v02.api.letsencrypt.org/directory", k)
	if err != nil {
		t.Error(err)
		return
	}
	accService, err := client.CreateAccount("hello@hello-example.com", true)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(accService)
	client.Persist("./a.json")
}
