package main

import (
	"github.com/zerocruft/keva"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {

	fileBytes, err := ioutil.ReadFile("test.properties")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileContent := string(fileBytes)

	props, err := keva.Parse(fileContent, "=")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(props)
}
