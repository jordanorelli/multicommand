package main

import (
	"fmt"
	"os"
	"strconv"
)

var addF bool // -f flag

var cmdAdd = &Command{
	UsageLine: "add [n0,n1,n2,...]",
	Short:     "adds numbers",
	Long: `
the add subcommand adds together successive strings of numbers.
`,
	Run: runAdd,
}

func runAdd(cmd *Command, args []string) {
	if addF {
		acc := 0.0
		for _, s := range args {
			v, err := strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
				os.Exit(3)
			}
			acc += v
		}
		fmt.Fprintln(os.Stdout, acc)
	} else {
		acc := 0
		for _, s := range args {
			v, err := strconv.Atoi(s)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
				os.Exit(3)
			}
			acc += v
		}
		fmt.Fprintln(os.Stdout, acc)
	}
}

func init() {
	cmdAdd.Flag.BoolVar(&addF, "f", false, "add floating-point numbers instead of integers")
	addCommand(cmdAdd)
}
