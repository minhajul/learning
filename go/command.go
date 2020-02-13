package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func execute() {
	output, err := exec.Command("ls", "-ltr").Output()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(output))
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can not run in windows")
	} else {
		execute()
	}
}