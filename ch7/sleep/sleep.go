package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period",
	1*time.Second,
	`Duration for which the program will sleep. Input might be specified with postfixes {h, s, ms, ns}`)

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
