package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	grepText = flag.String("grepText", "some text", "Text for grep")
	grepFile = flag.String("grepFile", "test.txt", "File for grep")
)

func main() {
	flag.Parse()
	bufFile, err := ioutil.ReadFile(string(*grepFile))
	if err != nil {
		fmt.Println("File or grep text not found!!!")
		return
	}
	fmt.Println(strings.Contains(string(bufFile), *grepText))
}
