package getting_started_go

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Name struct {
	firstName string
	lastName  string
}

func main() {
	path := ""
	fmt.Println("Please input file full Path(Such as D:\\name.txt):")
	fmt.Scanln(&path)
	var name []Name

	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	rd := bufio.NewReader(f)

	for {
		line, _, err := rd.ReadLine()

		if err != nil || io.EOF == err {
			break
		}
		tempName := strings.Split(string(line), " ")
		typeName := Name{
			fixName(tempName[0]),
			fixName(tempName[1]),
		}

		name = append(name, typeName)
	}
	for i := 0; i < len(name); i++ {
		fmt.Println(i + 1)
		fmt.Println("First Name:" + name[i].firstName)
		fmt.Println("Last Name:" + name[i].lastName)
	}
}

func fixName(buffer string) string {
	if len(buffer) > 20 {
		return string(buffer[0:20])
	}

	return buffer
}
