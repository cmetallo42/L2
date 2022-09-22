package main

import (
	"os"

	iinternal "calendar/internal"
	idatabase "calendar/internal/database"

	"github.com/sakirsensoy/genv"
	"github.com/sakirsensoy/genv/dotenv"
)

var (
	envfile string = ".env"
)

func main() {
	if len(os.Args) > 1 {
		envfile = os.Args[1]
	}

	dotenv.Load(envfile)

	configuration := iinternal.Configuration{
		Host:  genv.Key("HOST").String(),
		Database: idatabase.Configuration{
			DSN: genv.Key("DATABASE_DSN").String(),
		},
	}

	err := iinternal.Main(&configuration)
	if err != nil {
		panic(err)
	}
}
