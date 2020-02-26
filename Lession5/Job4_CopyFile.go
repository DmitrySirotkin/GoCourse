package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

var (
	copyFile = flag.String("copyFile", "test.txt", "File for copy")
)

func main() {
	flag.Parse()
	bufFile, err := ioutil.ReadFile(string(*copyFile))
	if err != nil {
		fmt.Println("File for copy not found!!!")
		return
	}
	err = ioutil.WriteFile("copy_"+string(*copyFile), bufFile, 755)
	if err != nil {
		fmt.Println("Can't write file!!!")
	} else {
		fmt.Println("Done!")
	}
}
