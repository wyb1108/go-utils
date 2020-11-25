package util

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func CheckErrorAndExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func CheckErrorAndPanic(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
