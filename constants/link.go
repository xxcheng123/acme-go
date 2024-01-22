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

func GetLinks(header http.Header) Links {
	hs := header.Values("Link")
	var links Links = make([]*Link, len(hs))
	for i, v := range hs {
		if link := parseLink(v); link != nil {
			links[i] = link
		}
	}
	return links
}

var linkReg = regexp.MustCompile(`<(.+?)>(?:;[^;]+)*?;\s*rel="(.+?)"`)

func parseLink(s string) *Link {
	result := linkReg.FindAllStringSubmatch(s, -1)
	if result == nil {
		return nil
	} else {
		return &Link{
			Url: result[0][1],
			Rel: result[0][2],
		}
	}
}
