package main

import (
	"github.com/zerocruft/keva"
	"fmt"
	"os"
)

func main() {

	props, err := keva.Open("test.properties")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(props)
}
