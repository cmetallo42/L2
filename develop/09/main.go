package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	URL := ""
	if len(os.Args) < 2 {
		fmt.Println("No URL. Setting default URL...")
		fmt.Println("Usage: go run main.go 'URL'")
		URL = "https://malinajs.github.io/docs"
	} else {
		URL = os.Args[1]
	}
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName := "index.html"
	err = ioutil.WriteFile(fileName, body, 0o644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Done")
}
