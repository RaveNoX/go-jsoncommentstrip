package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/RaveNoX/go-jsoncommentstrip"
)

func main() {
	b, err := ioutil.ReadAll(jsoncommentstrip.NewReader(os.Stdin))
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
	fmt.Print(string(b))
}
