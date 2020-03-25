package main

import "fmt"

func main() {
	test2()
}
func test2() {
	s := []int{0, 1, 2, 3, 4, 5}
	s1 := s[0:0]
	fmt.Println(s1) // return []

	s2 := s[0:1]
	fmt.Println(s2) // return [0]
}

func test1() {
	// panic: runtime error: slice bounds out of range [:50] with capacity 6
	s := []int{0, 1, 2, 3, 4, 5}
	s1 := s[:50]
	fmt.Println(s1)
}
