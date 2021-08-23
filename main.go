package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(errors.New("no args"))
	}

	in := os.Args[1]
	res, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		log.Fatalf("decoding: %v", err)
	}

	fmt.Println(string(res))
}
