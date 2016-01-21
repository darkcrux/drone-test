package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for _, i := range os.Args[1:] {
			fmt.Println(i)
		}
	}

}
