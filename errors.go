package main

import (
	"fmt"
	"errors"
)

func f1(arg int) (int, error) {
	if arg==42 {
		return -1, errors.New("can't work with 42 !")
	}
	return arg+3, nil
}

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		//return -1, &argError{arg, "can't work with it!"}
		e := argError{arg: arg, prob: "can't work with it!" }
		return -1, &e
	}
	return arg+3, nil
}

func main() {
	if r, e :=f1(42); e != nil {
		fmt.Println(e)
	}else{
		fmt.Println(r)
	}

	if r, e :=f2(42); e != nil {
		fmt.Println(e)
	}else{
		fmt.Println(r)
	}
	_, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }	
	
}