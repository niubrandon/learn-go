package knowledges

/*
strings and runes and bytes
*/


/*
nil map, length 0, 
read a nil map always returns the zero, 
write to a nil map cariable causes a panic
*/

// var nilMap map[string]int

/* map variable */
// totalWins := map[string]int{}

/* map literal
*/
// teams := map[string][]string{
// 	"Orcas": []string{"Fred", "Ralph", "Bijou"},
// 	"Lions": []string{"Sarah", "Peter", "Billie"},
// }

/*
map[keyType]valueType
the comma ok idiom:
deleting from Maps 
*/
// m := map[string]int{
// 	"hello": 5,
// 	"world": 0,
// }
// v, ok := m["hello"]
// fmt.Println(v, ok)

// v, ok := m["goodbye"]
// fmt.Println(v, ok)

// delete(m, "hello")

/*
map
*/


/*
using Maps as Sets

*/

m := map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
}

for i:=0; i < 3; i++ {
	fmt.Println("Loop", i)
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// iterating over strings
samples := []string{"hello", "apple_n!"}
for _, sample := range samples {
	for i, r := range sample {
		fmt.Println(i, r, string(r))
	}
	fmt.Println()
}
// 0 104 h
// 1 101 e
// 2 108 l
// 3 108 l
// 4 111 o

// 0 97 a
// 1 112 p
// 2 112 p
// 3 108 l
// 4 101 e
// 5 95 _
// 6 960 n
// 8 33 !
// (the first column skips the number 7, the value at postion 6 is 960)
// it is far larger than what can fit in a byte, strings were made out of bytes
// it iterates over the runes, not the bytes. when a for - range loop encounters a multibyte rune in a string, 
// it converts the UTF-8 representation into a signle 32-bit number and assigns it to the value

/*
for range value is a copy
*/

evenVals := []int{2, 4, 6, 8, 10, 12}
for _, v := range evenVals {
	v *= 2
}
fmt.Println(evenVals)

// [2 4 6 8 10 12]