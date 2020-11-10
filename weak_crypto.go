package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func makeMD5Hash() {
	for _, arg := range os.Args {
		fmt.Printf("%x - %s\n", md5.Sum([]byte(arg)), arg)
	}
}
