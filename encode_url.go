package idev

import (
	"net/url"
)

func EncodeUrl(s string) string {
	return url.PathEscape(s)
}

func main() {

}
