package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gofrs/uuid"
	"github.com/peterbourgon/ff/v3"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fs := flag.NewFlagSet("genid", flag.ExitOnError)
	var (
		noNewline = fs.Bool("n", false, "Do not print the trailing newline character.")
	)

	ff.Parse(fs, args,
		ff.WithEnvVarPrefix("GENID"),
	)

	var lineEnding string
	if *noNewline == false {
		lineEnding = "\n"
	}

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	fmt.Printf("%s%s", id, lineEnding)

	return nil
}
