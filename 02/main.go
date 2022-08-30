package main

import (
	"fmt"
	"os"
	"strconv"
)

func parser(str string) string {
	newStr := ""
	flag := false
	for i, r := range str {
		n, err := strconv.Atoi(string(r))
		if err == nil && flag {
			for j := 1; j < n; j++ {
				newStr += string(str[i-1])
			}
			flag = false
		} else if err != nil {
			newStr += string(r)
			flag = true
		} else if !flag {
			return "Incorrent string"
		}
	}
	return newStr
}

func main() {
	str := "a4bc2d5e"
	if len(os.Args) > 1 {
		str = os.Args[1]
	}

	fmt.Println(parser(str))
}
