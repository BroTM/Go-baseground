package main

import "fmt"

type Pet interface {
	shout() string
}

type Dog struct {

}

type Cut struct {

}

func (d *Dog) shout() string{
	return "Wote Wote"
}
func (c Cut)  shout() string{
	return "Nhaung Nhaung"
}

func feature(p Pet)  {
	fmt.Println(p.shout())
}

func main()  {
	d := new(Dog) //use new keyword because of pointer receiver
	c := Cut{}

	feature(d)
	feature(c)

	fmt.Println(d.shout())
	fmt.Println(c.shout())
}