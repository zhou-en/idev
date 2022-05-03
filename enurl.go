package idev

import (
	"fmt"
	"net/url"
	"os"
)

func EncodeUrl(s string) string {
	return url.PathEscape(s)
}
