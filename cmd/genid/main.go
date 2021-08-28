package main

import (
	"fmt"
	"os"

	"github.com/gofrs/uuid"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	fmt.Println(id)

	return nil
}
