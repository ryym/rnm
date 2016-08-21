package main

import (
	"fmt"
	"github.com/ryym/rnm"
	"os"
)

func main() {
	opts := rnm.Option{}

	results, err := rnm.Exec(opts)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(results)
}
