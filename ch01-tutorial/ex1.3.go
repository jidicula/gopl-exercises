package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Printf("=Explicit For loop:\n")

	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.2fs elapsed\n\n", time.Since(start).Seconds())

	fmt.Printf("=Range For loop:\n")
	start = time.Now()
	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.2fs elapsed\n\n", time.Since(start).Seconds())

	fmt.Printf("=strings.Join():\n")
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.2fs elapsed\n\n", time.Since(start).Seconds())
}
