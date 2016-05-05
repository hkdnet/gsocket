package lib

import (
	"fmt"
	"os"
)

func DealError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: erro: %s", err.Error())
		os.Exit(1)
	}
}
