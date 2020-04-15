package main

import (
	"flag"
)

func main() {
	var flagValue int
	flag.IntVar(&flagValue, "flagname", 1234, "help message for flagname")
}
