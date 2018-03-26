package main

import "fmt"

//Slice manipulation
func main() {
	s := []int{1,2,3,4,5,6,7,8,9}
	printSlice(s)
	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)
	//Give it 4 lenght
	s = s[:4]
	printSlice(s)
	//Slice first two
	s = s[2:]
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}