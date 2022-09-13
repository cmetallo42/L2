package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var files []string
	flag := false
	var pathPlusFile string
	path := "examples/"
	var substr string
	keys := make(map[rune]int)
	sliceOfStrings := make([]string, 0, 1)

	for i, arg := range os.Args {
		if arg[0] == '-' {
			for _, r := range arg {
				if r == 'A' || r == 'B' || r == 'C' {
					nextI, err := strconv.Atoi(os.Args[i+1])
					if err != nil {
						panic(err)
					}
					keys[r] = nextI
					flag = true
				} else {
					keys[r] = 0
				}
			}
		} else if i < len(os.Args)-1 && i > 0 && arg != "examples" && arg != "main.go" {
			if !flag {
				if substr == "" {
					substr += arg
				} else {
					substr += " " + arg
				}
			} else {
				flag = false
			}
		}
	}

	if substr == "" {
		fmt.Println("No Substring")
		return
	}

	delete(keys, '-')
	_, keyA := keys['A']
	_, keyB := keys['B']
	_, keyC := keys['C']
	// _, keyc := keys['c']
	// _, keyi := keys['i']
	// _, keyv := keys['v']
	// _, keyF := keys['F']
	_, keyn := keys['n']

	if strings.Contains(os.Args[len(os.Args)-1], "main.go") || strings.Contains(os.Args[len(os.Args)-1], "*.txt") {
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		fileInfo, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			panic(err)
		}

		for _, file := range fileInfo {
			files = append(files, file.Name())
		}

		for _, file := range files {
			tmp, err := os.ReadFile(path + file)
			if err != nil {
				panic(err)
			}
			pathPlusFile = path + file + ":"
			tmpStr := strings.Split(string(tmp), "\n")
			for i, t := range tmpStr {
				iStr := strconv.Itoa(i + 1)
				if strings.Contains(t, substr) {
					if keyn {
						pathPlusFile = path + file + ":" + iStr + ":"
					}
					if keyA {
						aInt := keys['A']
						for j := 0; j <= aInt; j++ {
							if i+j > len(tmpStr)-1 {
								break
							}
							if keyn {
								pathPlusFile = path + file + ":" + strconv.Itoa(i+j) + ":"
							}
							sliceOfStrings = append(sliceOfStrings, pathPlusFile+tmpStr[i+j])
						}
					}
					if keyB {
						bInt := keys['B']
						for j := bInt; j >= 0; j-- {
							if i-j < 0 {
								break
							}
							if keyn {
								pathPlusFile = path + file + ":" + strconv.Itoa(i-j) + ":"
							}
							sliceOfStrings = append(sliceOfStrings, pathPlusFile+tmpStr[i-j])
						}
					}
					if keyC {
						cInt := keys['C']
						for j := 0 - cInt; j <= cInt; j++ {
							if i+j < 0 || i+j > len(tmpStr)-1 {
								break
							}
							if keyn {
								pathPlusFile = path + file + ":" + strconv.Itoa(i+j) + ":"
							}
							sliceOfStrings = append(sliceOfStrings, pathPlusFile+tmpStr[i+j])
						}
					}
				}
			}
		}
	}
	for _, s := range sliceOfStrings {
		fmt.Println(s)
	}
}
