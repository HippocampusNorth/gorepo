package main

import (
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
)
//fetchurl http://www.baidu.com
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if nil != err {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if nil != err {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
	}
}
