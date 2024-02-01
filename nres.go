package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Resolve Nessus URL from Report
	args := os.Args[1:]
	uri := args[0]
	fmt.Printf("%s => %s\n", args[0], NessusURLResolver(uri))
}

/*
##
	This resolver uses api.tenable.com not nessus.org this action seems to take less time with a halt on the redirect getting the URI from the Location header.
##
*/
func NessusURLResolver(url string) string {
	uri := strings.Replace(url, "http://www.nessus.org/", "https://api.tenable.com/v1/", -1)
	Header := map[string][]string{
		"Host":            {"api.tenable.com"},
		"Accept":          {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"},
		"Accept-Language": {"en-GB,en;q=0.5"},
		"Accept-Encoding": {"gzip, deflate"},
		"Content-Type":    {"text/html"},
		"Content-Length":  {"0"},
		"Origin":          {"https://api.tenable.com"},
	}
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header = Header
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	location := response.Header["Location"][0]
	defer response.Body.Close()
	return location
}
