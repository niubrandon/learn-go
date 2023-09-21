package knowledges

import (
	"fmt"
	"strconv"
	"sort"
)
/*
function as value
*/

func add(i int, j int) int { return i + j }

func sub(i int, j int) int { return i - j }

func mul(i int, j int) int { return i * j }

func div(i int, j int) int { return i / j }

// example build a map using fcunction as value
var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

//
expressions := [][]string{
	[]string{"2", "+", "3"},
	[]string{"2", "-", "3"},
	[]string{"2", "*", "3"},
	[]string{"2", "/", "3"},
	[]string{"2", "%", "3"},
	[]string{"two", "+", "three"},
	[]string{"5"},
}

for _, expression := range expressions {
	if len(expression) != 3 {
		fmt.Println("invalid expression:", expression)
		continue
	}
	// strconv.Atoi function i nthe standard library to 
	// convert a string to an int, second value return by this function is an error
	p1, err := strconv.Atoi(expression[0])
	if err != nil {
		fmt.Println(err)
		continue
	}
	op := expression[1]
	// 
	opFunc, ok := opMap[op]
	if !ok {
		fmt.Println("unsupported operator:", op)
		continue
	}
	p2, err := strconv.Atoi(expression[2])
	if err != nil {
		fmt.Println(err)
		continue
	}
	result := opFunc(p1, p2)
	fmt.Println(result)
}

// function type declarations

type opFuncType func(int, int) int

var opMap = map[string]opFuncType {
	// same as before
}

/*
Anonymous function
*/

for i := 0; i < 5; i++ {
	func(j int) {
		fmt.Println("printing", j, "from inside of an anonymous function")
	}(i)
}

// printing 0 from inside of an anonymous function
// printing 1 from inside of an anonymous function
// printing 2 from inside of an anonymous function
// printing 3 from inside of an anonymous function
// printing 4 from inside of an anonymous function

/*
Closures: functions declared inside of functions are special; they are closures
functions declared inside of functions are able to access and modify variables declared in the outer function
*/


/*
Passing functions as parameters
*/
type Person struct {
	FirstName string
	LastName  string
	Afe 			int
}

people := []Person{
	{"Pat", "Patterson", 37},
	{"Tracy", "Bobbert", 23},
	{"Fred", "Fredson", 18},
}

fmt.Println(people)

// sort by last name
sort.Slice(people, func(i int, j int) bool {
	return people[i].LastName < people[j].LastName
})
fmt.Println(people)

// sort by age
sort.Slice(people, func(i int, j int) bool {
	return people[i].Age < people[j].Age
})
fmt.Println(people)

/*
Returning functions from functions
*/
// you can return a closure from a function
func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

twoBase := makeMult(2)
threeBase := makeMult(3)
for i := 0; i < 3; i++ {
	fmt.Println(twoBase(2), threeBase(3))
}
// 0 0
// 2 3
// 4 6

/*
Defer
*/