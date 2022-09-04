package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var files []string
	path := "examples"
	var substr string
	keys := make(map[string]struct{})
	sliceOfStrings := make([]string, 0, 1)

	if len(os.Args) > 1 {
		substr = os.Args[1]
	}
	if substr == "" {
		fmt.Println("No Substring")
		return
	}

	for _, args := range os.Args {
		if args[0] == '-' {
			for _, r := range args {
				keys[string(r)] = struct{}{}
			}
		}
	}
	delete(keys, "-")
	//    _, keyA := keys["A"]
	// _, keyB := keys["B"]
	// _, keyC := keys["C"]
	// _, keyc := keys["c"]
	// _, keyi := keys["i"]
	// _, keyv := keys["v"]
	// _, keyF := keys["F"]
	_, keyn := keys["n"]

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
		tmp, err := os.ReadFile(path + "/" + file)
		if err != nil {
			panic(err)
		}
		tmpStr := strings.Split(string(tmp), "\n")
		for i, t := range tmpStr {
			iStr := strconv.Itoa(i + 1)
			if keyn && strings.Contains(string(t), substr) {
				sliceOfStrings = append(sliceOfStrings, path+"/"+file+":"+iStr+":"+string(t)+"\n")
			} else if !keyn && strings.Contains(string(t), substr) {
				sliceOfStrings = append(sliceOfStrings, string(t)+"\n")
			}
		}

	}

	fmt.Println(sliceOfStrings)
}
