package constants

import (
	"net/http"
	"testing"
)

/**
*	@Author: xxcheng
*	@Email developer@xxcheng.cn
*	@Date: 2024/1/19
 */

func TestGetLinks(t *testing.T) {
	testCases := []*struct {
		Name   string
		Header http.Header
		Links  Links
	}{
		{
			Name: "pass",
			Header: http.Header{
				"Link": []string{
					`<https://example.com/acme/directory>;rel="index"`,
					`<https://example.com/acme/orders/rzGoeA?cursor=2>;rel="next"`,
				},
			},
			Links: Links{
				{
					Url: "https://example.com/acme/directory",
					Rel: "index",
				}, {
					Url: "https://example.com/acme/orders/rzGoeA?cursor=2",
					Rel: "next",
				},
			},
		}, {
			Name:   "empty",
			Header: http.Header{},
			Links:  Links{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			links := GetLinks(tc.Header)
			if len(links) != len(tc.Links) {
				t.Errorf("GetLinks() = %v, want %v", links, tc.Links)
			}
		})
	}
}
