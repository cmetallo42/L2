package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	keys := make(map[rune]string)
	strs := make([]string, 0, 1)
	res := make([]string, 0, 1)
	flag := false
	flag2 := false

	for i, arg := range os.Args {
		if flag2 {
			flag2 = !flag2
			continue
		}
		if i == 0 {
			continue
		}
		if arg[0] == '-' {
			for j, r := range arg {
				if j == 0 {
					continue
				}
				if r == 's' {
					keys['s'] = ""
				} else {
					if i+1 > len(os.Args) {
						panic("Out of range")
					}
					keys[r] = string(os.Args[i+1])
					flag2 = true
				}

			}
		} else {
			strs = append(strs, arg)
		}
	}

	keyFvalue, keyF := keys['f']
	keyDvalue, keyD := keys['d']
	_, keyS := keys['s']

	if !keyF {
		fmt.Println("You must specify a list of fields")
		return
	}
	keyFint, err := strconv.Atoi(keyFvalue)
	if err != nil {
		panic(err)
	}

	if !keyD {
		keyDvalue = "\t"
	}

	for _, sepStr := range strs {
		flag = false
		for i, str := range strings.Split(sepStr, keyDvalue) {
			if keyFint == i+1 {
				res = append(res, str, "\n")
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		if !keyS {
			res = append(res, sepStr, "\n")
		}
	}
	fmt.Println(res)
}
