package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func event_listener() {
	var tmp string
	var arr []string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type help or h for help")
	for {
		fmt.Print("[cmetallo@fedora]$ ")
		scanner.Scan()
		tmp = scanner.Text()
		arr = strings.Split(tmp, " ")
		if arr[0] == "exit" || arr[0] == "q" || arr[0] == "quit" {
			fmt.Println("Exit...")
			return
		}
		if arr[0] == "help" || arr[0] == "h" {
			help()
			continue
		}
		if arr[0] == "cd" && len(arr) == 2 {
			os.Chdir(arr[1])
		} else {
			execute(arr)
			flag = false
			fileName := "tmp.txt"
			_ = ioutil.WriteFile(fileName, []byte(""), 0o644)
		}

	}
}

var flag bool

// cat some.txt | grep alex
func execute(arr []string) {
	j := 0
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == "|" {
			flag = true
			execute(arr[:i])
			args := append(arr[i+2:], "tmp.txt")
			cmd := exec.Command(arr[i+1], args...)
			output, _ := cmd.CombinedOutput()
			fmt.Println(string(output))
			j = i
		}
		if i == 0 && j == 0 {
			cmd := exec.Command(arr[0], arr[1:]...)
			output, _ := cmd.CombinedOutput()
			fileName := "tmp.txt"
			err := ioutil.WriteFile(fileName, output, 0o644)
			if err != nil {
				log.Fatal(err)
			}
			if !flag {
				fmt.Println(string(output))
			}
		}

	}
}

func help() {
	fmt.Println("Type command to execute it.\nType exit | q | quit to exit")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		event_listener()
	}()
	wg.Wait()
}
