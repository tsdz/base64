package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {
	in := os.Args[1]
	res, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		log.Fatalf("decoding: %v", err)
	}

	fmt.Println(string(res))
}
