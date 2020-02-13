package main

import (
	"os"
	"fmt"
	"runtime"
)

func main()  {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
		fmt.Println(os.Hostname())
		fmt.Println(os.Geteuid())
		fmt.Println(runtime.GOOS)
	}else {
		fmt.Println("Hello, I am gopher!")
	}
}
