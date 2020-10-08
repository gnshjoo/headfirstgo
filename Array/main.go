package main

import (
	"fmt"
	"time"
)

func main() {
	var notes [7]string
	notes[0]= "do"
	notes[1]= "re"
	notes[2]= "mi"

	notes2 := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	primes := [5]int{2, 3, 5, 7, 11}

	text := [3]string{
		"This is a series of long strings",
		"which would be awkward to palce",
		"together on a single line",
	}

	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(4532540000, 0)
	dates[2] = time.Unix(1251234000, 0)

	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(dates[1])
}
