package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
)

func main() {

	if t, err := ntp.Time("pool.ntp.org"); err == nil {
		fmt.Println(t)
	} else {
		l := log.New(os.Stderr, "[ERROR]", 0)
		l.Println(err)
		os.Exit(2)
	}
}
