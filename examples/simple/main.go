package main

import (
	"fmt"

	"github.com/erdaltsksn/jerr"
)

func main() {
	err := someFunc()
	if err != nil {
		wrapped := jerr.Wrap(err, "Message about error")
		fmt.Println(wrapped.Error())
	}
}

func someFunc() error {
	return jerr.New("nope")
}
