package main

import (
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SliceOfSliceOfString struct {
	values  [][]string
	column  int
	reverse bool
	number  bool
}

func less(firsts string, seconds string) (bool, error) {
	first := big.Float{}
	_, _, one := first.Parse(firsts, 0)

	second := big.Float{}
	_, _, two := second.Parse(seconds, 0)

	if one != nil && two != nil {
		return false, two
	} else if one != nil {
		return true, nil
	} else if two != nil {
		return false, nil
	}

	return first.Cmp(&second) < 0, nil
}

func (x *SliceOfSliceOfString) Less(i, j int) (l bool) {
	if x.reverse {
		defer func() { l = !l }()
	}

	first := x.values[i][0]
	second := x.values[j][0]

	fe := len(x.values[i]) > x.column
	se := len(x.values[j]) > x.column

	if fe && se {
		first = x.values[i][x.column]
		second = x.values[j][x.column]
	} else if fe {
		return
	} else if se {
		l = true
		return
	}

	if x.number {
		n, err := less(first, second)
		if err == nil {
			l = n
			return
		}
	}

	l = first < second
	return
}

func (x *SliceOfSliceOfString) String() (s string) {
	for i := range x.values {
		s += strings.Join(x.values[i], " ") + "\n"
	}

	return
}

func (x *SliceOfSliceOfString) Len() int      { return len(x.values) }
func (x *SliceOfSliceOfString) Swap(i, j int) { x.values[i], x.values[j] = x.values[j], x.values[i] }
func (x *SliceOfSliceOfString) Sort()         { sort.Sort(x) }

func main() {
	keys := make(map[string]struct{})

	unsorted := SliceOfSliceOfString{
		values: make([][]string, 0, 2),
		column: 0,
	}

	for _, args := range os.Args {
		if _, err := strconv.Atoi(args); err == nil {
			unsorted.column, _ = strconv.Atoi(args)
			unsorted.column -= 1
		}
		if args[0] == '-' {
			for _, r := range args {
				keys[strings.ToLower(string(r))] = struct{}{}
			}
		}
	}

	file, _ := os.ReadFile("some.txt")
	strs := string(file)

	_, keyK := keys["k"]
	_, keyN := keys["n"]
	_, keyR := keys["r"]
	_, keyU := keys["u"]

	ss := strings.Split(strs, "\n")

	if keyU {
		exists := make(map[string]struct{}, len(ss))

		for i := range ss { // бьём файл на строки
			_, ok := exists[ss[i]]
			if ok {
				continue
			}

			unsorted.values = append(unsorted.values, strings.Split(ss[i], " ")) // записываем массивы слов как строки
			exists[ss[i]] = struct{}{}
		}
	} else {
		for i := range ss { // бьём файл на строки
			unsorted.values = append(unsorted.values, strings.Split(ss[i], " ")) // записываем массивы слов как строки
		}
	}

	if keyK {
		if unsorted.column < 0 {
			fmt.Println("Flag -k needs value > 0")
			return
		}
	}

	if keyR {
		unsorted.reverse = true
	}

	if keyN {
		unsorted.number = true
	}

	sort.Sort(&unsorted)

	fmt.Println(unsorted.String())
}
