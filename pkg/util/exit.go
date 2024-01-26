package util

import (
	"fmt"
	"os"
)

func ExitOnError(err error) {
	if err == nil {
		return
	}

	fmt.Println(err)
	os.Exit(1)
}
