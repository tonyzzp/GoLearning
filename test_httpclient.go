package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := http.Client{Transport: t}
	resp, e := client.Get("https://localhost")
	if e == nil {
		content, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(content))
	} else {
		fmt.Println(e)
	}
}
