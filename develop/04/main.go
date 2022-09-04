package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

func Anagramma(words []string) (m map[string][]string) {
	mtmp := make(map[string][]string)

root:
	for i := range words {
		if !utf8.ValidString(words[i]) {
			fmt.Println("UTF-8 check failed")

			return
		}

		original := strings.ToLower(words[i])

		r := []rune(original)

		sort.Slice(r, func(i int, j int) bool { return r[i] < r[j] })

		s := string(r)

		value, ok := mtmp[s]

		if ok {
			for j := range value {
				if value[j] == original {
					continue root
				}
			}
		}

		mtmp[s] = append(mtmp[s], original)
	}

	m = make(map[string][]string)

	for _, value := range mtmp {
		sort.Strings(value)
		m[value[0]] = value[1:]
	}

	return
}

func main() {
	words := []string{"Ромашка", "Мошкара", "Пятак", "Тест", "Пятка", "Столик", "Тяпка", "Слиток", "Мошкара", "Листок", "Материк", "Тяпка", "Тяпка", "Торнадо", "Метрика", "Керамит", "Тяпка", "Термика", "Ромашка", "Ротонда"}
	m := Anagramma(words)

	fmt.Println(m)
}
