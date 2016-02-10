package main

import (
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/royvandewater/etcdsync/keyvalue"
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

	log.Panicf("ERROR(%v):\n\n%v", msg, err)
}

func printKeyValuePairs(useTable bool, keyValues []keyvalue.KeyValue) {
	if useTable {
		printTable(keyValues)
		return
	}

	printPlain(keyValues)
}

func printPlain(keyValues []keyvalue.KeyValue) {
	for _, keyValue := range keyValues {
		fmt.Println(keyValue)
	}
}

func printTable(keyValues []keyvalue.KeyValue) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key", "Value"})

	for _, keyValue := range keyValues {
		table.Append([]string{keyValue.Key, keyValue.Value})
	}
	table.Render()
}
