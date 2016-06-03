// url
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Printf("no url input is found \n")
		os.Exit(1)
	}
	for _, url := range os.Args[1:] {

		if strings.HasPrefix(url, "http://") == false {
			url = strings.Join([]string{"http://", url}, "")
		}

		content, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(content.Body)
		content.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
