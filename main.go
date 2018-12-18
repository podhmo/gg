package main

import (
	"fmt"
	"log"
	"os"

	"github.com/podhmo/gg/resolve"
)

func main() {
	if len(os.Args) <= 2 {
		usage(1)
	}

	if err := run(os.Args[1], os.Args[2:]); err != nil {
		log.Fatalf("%+v", err)
	}
}

func usage(code int) {
	fmt.Print(`
usage: gg <commands> [arguments]
The commands are:

	resolve -- resolve package path from filepath
`)
	os.Exit(code)
}

func run(cmd string, args []string) error {
	switch cmd {
	case "resolve":
		return resolve.Main(args)
	default:
		usage(1)
	}
	return nil
}
