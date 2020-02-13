package main

import "fmt"

func anonymous() func() int  {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main(){

	nextInt := anonymous()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}