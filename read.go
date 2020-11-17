package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname [20]rune
	lname [20]rune
}

var sliceOfName = make([]Name, 0)

func main() {
	fmt.Println("enter the name of the text file you want to open:")

	var fileName string
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n := Name{}
		line := scanner.Text()
		actualLine := strings.Split(line, " ")

		fname := actualLine[0]
		fullRuneFName := []rune(fname)
		for i, rune := range fullRuneFName {
			if i > 19 {
				break
			}
			n.fname[i] = rune
		}

		lname := actualLine[1]
		fullRuneLName := []rune(lname)
		for i, rune := range fullRuneLName {
			if i > 19 {
				break
			}
			n.lname[i] = rune
		}

		sliceOfName = append(sliceOfName, n)
	}

	for _, n := range sliceOfName {
		fmt.Println(string(n.fname[:20]), string(n.lname[:20]))
	}

}
