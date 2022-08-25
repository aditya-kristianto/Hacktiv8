package main

import (
	"fmt"
	"strings"
)

type Person struct{ name string }

func main() {
	contohClosure := func(persons []*Person) {
		// cetak data-data dari list
		for i, v := range persons {
			fmt.Printf("%d. %s%s\n", i+1, strings.ToUpper(v.name[0:1]), v.name[1:])
		}
	}

	var friends = []string{"clara", "fiqri", "medy", "lutfi", "raiza", "ian", "umar", "fazry", "pandi", "tantut"}

	var persons []*Person

	for _, v := range friends {
		var person = &Person{
			name: v,
		}

		persons = append(persons, person)
	}

	contohClosure(persons)
}
