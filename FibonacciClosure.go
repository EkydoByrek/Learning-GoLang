package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f1 ,f0:= 1,0
	return func() int{
		temp:=f0
		f1,f0=f1+temp,f1
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
