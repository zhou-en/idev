package main

import (
	"fmt"
	"net/url"
	"os"
	"flag"
)

func EncodeUrl(s string) string {
	return url.PathEscape(s)
}

func main() {
	if len(os.Args) == 2 {
		url := os.Args[1]
		fmt.Println(EncodeUrl(url))
	}
}
