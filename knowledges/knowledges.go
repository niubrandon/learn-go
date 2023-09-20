package basics

import (
	"fmt"
	"time"
)

// statically typed lanugages with both built-in types and user-defined types
// user-defined types: have underlying type of the struct literal that follows
type Person struct {
	FirstName string
	LastName string
	Age int
}
// abstract type: specifies what a type should do, but not how it is done
// concrete type: specifies what and how
// concrete type: use any primitive type or compount type literal
type Score int
type Converter func(string)Score
type TeamScores map[string]Score

type Counter struct {
	total       int
	lastUpdated time.Time
}

// methods: like function declarations, with one additional the receiver specification

/* Section
Pointer Receivers and Value Receivers with Methods
*/
// receiver specification: pointer receiver
// 1. when method modifies the receiver,
// 2. when method needs to handle nil instances
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}
// receiver specification: value receiver
// if the method doesn't modify the recievers
func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}
// var c *Counter
//	fmt.Println(c.String())
//	c.Increment()
//	fmt.Println(c.String())

/* Section
Pointer Receivers and Value Receivers with Functions
*/
func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}
// var c *Counter
// doUpdateWrong(c)
// fmt.Println("in main:", c.String())
// doUpdateRight(&c)
// fmt.Println("in main:", c.String())

/* Section
nil instances
*/

type IntTree struct {
	val int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}

	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

// var it *IntTree
// it = it.Insert(5)
// it = it.Insert(3)
// it = it.Insert(10)
// it = it.Insert(2)
// fmt.Println(it.Contains(2)) // true
// fmt.Println(it.Contains(12)) // false