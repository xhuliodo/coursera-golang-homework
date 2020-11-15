package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Name struct is name of lines
// TODO: modify name struct with strings of only 20 charaters in length
type Name struct {
	//  fname [20]rune
	fname string
	lname string
}

var sliceOfName = make([]Name, 0)

func main() {
	var fileName string
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		actualLine := strings.Split(line, " ")
		if err != nil {
			break
		}
		fname := actualLine[0]
		lname := strings.ReplaceAll(actualLine[1], "\r\n", "")
		sliceOfName = append(sliceOfName, Name{fname, lname})
	}

	for _, n := range sliceOfName {
		fmt.Println("First name: " + n.fname + " Last name: " + n.lname)
	}

}
