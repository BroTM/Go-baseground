package main

import (
	"fmt"
	"math"
)

const s string= "constant"
func main()  {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
    fmt.Println(d)

	var sth int64= int64(d)
	fmt.Println(sth)

	fmt.Println(math.Sin(n))
}