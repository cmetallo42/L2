package main

import (
	"fmt"

	"github.com/beevik/ntp"
)

func main() {
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}
