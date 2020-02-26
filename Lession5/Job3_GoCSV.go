package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	csvFile := [][]string{
		{"User_name", "login", "passwd"},
		{"Dima", "rusdsi", "qwed13"},
		{"Ivan", "rusimn", "ddfs"},
		{"Valera", "rusvva", "1123"},
	}
	bufCsvFile := csv.NewWriter(os.Stdout)
	bufCsvFile.WriteAll(csvFile)

	if err := bufCsvFile.Error(); err != nil {
		fmt.Println("Can't write file!")
	}
}
