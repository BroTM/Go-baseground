package main

import "fmt"

func main()  {
	/**Maps are Go's builtin assosiative data type
	* (sometimes called hashes or dicts in other languages).
	* make(map[key-type]val-type).
	* name[key] = val syntax.
	*/
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
    fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	// delete(map, "key")
	delete(m, "k2")
	fmt.Println("map:", m)


	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	m2 := make(map[string]string)
	m2["k1"] = "Something"
	m2["k2"] = "Something"
	m2["k3"] =""

	_, pr := m2["k3"]
	fmt.Println("prs :", pr);

	//slice
	s := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("slice", s)

	//map
	n := map[string]int{"foo": 1, "bar": 2, "gg": 4, "a": 0}
    fmt.Println("map:", n)
}