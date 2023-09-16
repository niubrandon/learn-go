package main

import (
	"fmt"

	_ b "github.com/niubrandon/learn-go/basics"
)

func Main() {
	var it *b.IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2)) // true
	fmt.Println(it.Contains(12)) // false
}