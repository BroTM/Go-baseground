package main

import ("fmt")

func main() {
	array := [3]int{1,3,5}
	

	//slice
	sl1 := make([]int, 2)
	sl2 := []int{1,2,3,4,5,6,7}
	sl3 := [...]int{1,2,4,5,6,6,7,7,8,9,0,2}

	//map
	m1 := make(map[string]int)
	m2 := map[string]int
	m3 := map[string]int{"a" : 1, "b" : 2}

	//channel
	ch := make(chan string)// unbuffered
	ch1 := make(chan string, 2) //buffered (first in first out)

}