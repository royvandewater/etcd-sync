package main

import (
	"log"
	"os"
)

// ExitWithHelp prints message and dies with exit code 1
func ExitWithHelp(msg string) {
	log.Print(msg)
	os.Exit(1)
}

// PanicIfError prints error and dies if error is non nil
func PanicIfError(msg string, err error) {
	if err == nil {
		return
	}

	log.Panicf("ERROR(%v): %v", msg, err)
}
