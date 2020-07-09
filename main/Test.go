package main

import "fmt"

func main() {
	m := 10000.00
	for i := 0; i < 120; i++ {
		m = m + (m * 0.04)
	}
	fmt.Println(m)
}
