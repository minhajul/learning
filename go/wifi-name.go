package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"runtime"
)

const osxCmd = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"
const osxArgs = "-I"

func main() {
	platform := runtime.GOOS
	if platform == "darwin" {
		fmt.Println(OSX())
	} else {
		fmt.Println("This program only for OSX!")
	}
}

func OSX() string {
	cmd := exec.Command(osxCmd, osxArgs)

	stdout, err := cmd.StdoutPipe()
	panicIf(err)

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	var str string

	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += string(b) + "\n"
	}

	name := regexp.MustCompile(`s*SSID: (.+)s*`)

	wifiName := name.FindAllStringSubmatch(str, -1)

	if len(wifiName) <= 1 {
		return "Could not find wifi"
	}

	return "Wifi name is " + wifiName[1][1]
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
