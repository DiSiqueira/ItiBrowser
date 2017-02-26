package itibrowser

import (
	"github.com/mssola/user_agent"
	"net/http"
	"strings"
)

// BrowserMatcher Store the schema to match with the request Path
type BrowserMatcher struct {
	browserList []string
}

// New is the constructor to ItiBrowser
func New(browser ...string) *BrowserMatcher {
	for i := range browser {
		browser[i] = strings.ToLower(browser[i])
	}
	return &BrowserMatcher{
		browserList: browser,
	}
}

// Match returns if the request can be handled by this Route.
func (b *BrowserMatcher) Match(req *http.Request) bool {
	ua := user_agent.New(req.UserAgent())
	name, _ := ua.Browser()
	name = strings.ToLower(name)
	for _, v := range b.browserList {
		if name == v {
			return true
		}
	}

	return false
}
