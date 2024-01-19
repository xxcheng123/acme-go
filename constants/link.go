package constants

import (
	"net/http"
	"regexp"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/19 17:47
 */

// Link ACME Server Response Header Link Fields
type Link struct {
	Url string
	Rel string
}
type Links []*Link

func getLinks(header http.Header) Links {
	hs := header.Get("Link")
	var links Links = make([]*Link, len(hs))
	for i, v := range hs {
	}
	panic("")
}

var linkReg = regexp.MustCompile(`<(.+?)>(?:;[^;]+)*?;\s*rel="(.+?)"`)

func parseLink(s string) *Link {
	linkReg.FindAllStringSubmatch(s, -1)
}
