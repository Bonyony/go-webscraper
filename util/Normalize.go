package util

import (
	"strings"
)

// ensures a URL contains "https://"
func NormalizeURL(domain string) string {
	if !strings.HasPrefix(domain, "http://") && !strings.HasPrefix(domain, "https://") {
		return "https://" + domain
	}
	return domain
}
