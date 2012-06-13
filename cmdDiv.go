package main

import (
	"fmt"
	"os"
	"strconv"
)

var divF bool // -f flag

var cmdDiv = &Command{
	UsageLine: "div [n0,n1,n2,...]",
	Short:     "divides numbers",
	Long: `
the div subcommand divides successive strings of numbers.
`,
	Run: runDiv,
}

func runDiv(cmd *Command, args []string) {
	if divF {
		acc := 1.0
		for _, s := range args {
			v, err := strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
				os.Exit(3)
			}
			acc /= v
		}
		fmt.Fprintln(os.Stdout, acc)
	} else {
		acc := 1
		for _, s := range args {
			v, err := strconv.Atoi(s)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
				os.Exit(3)
			}
			acc /= v
		}
		fmt.Fprintln(os.Stdout, acc)
	}
}

func init() {
	cmdDiv.Flag.BoolVar(&divF, "f", false, "")
	addCommand(cmdDiv)
}
