package main

import "fmt"

func main()  {
	// using make builtin fun
	s := make([]string, 3)
	fmt.Println("emp:", s);

	s[0] = "abcd"
    s[1] = "b"
    s[2] = "c"

	fmt.Println("set:", s)
    fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e","f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
    copy(c,s)
	c =append(c, "g");
    fmt.Println("cpy:", c)

	l := s[2:5]
    fmt.Println("sl1:", l)

	//0 to 5
	l = s[:5]
    fmt.Println("sl2:", l)

	l = s[2:]
    fmt.Println("sl3:", l)

	// declaration like array
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 4)
	fmt.Println(twoD)
    for i := 0; i < 4; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

	i := make([]int, 3);

	fmt.Println(i);

	// size by compiler
	i2 := [...]int{1,2,3,4,5,6,7,2,0,34,99}

	fmt.Println(i2);

	// func make([]T, len, cap) []T

	var vals []int
	for i := 0; i < 5; i++ {
		vals = append(vals, i)
		fmt.Println("The length of our slice is:", len(vals))
		fmt.Println("The capacity of our slice is:", cap(vals))
	}
}

