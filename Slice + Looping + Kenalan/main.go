package main

import (
	"fmt"
	"strings"
)

func main() {
	var friends = []string{"clara", "fiqri", "medy", "lutfi", "raiza", "ian", "umar", "fazry", "pandi", "tantut"}

	for i, v := range friends {
		fmt.Printf("%d. %s%s\n", i+1, strings.ToUpper(v[0:1]), v[1:])
	}
}
