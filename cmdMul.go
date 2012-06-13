package main

import (
	"fmt"
	"os"
	"strconv"
)

var mulF bool // -f flag

var cmdMul = &Command{
	UsageLine: "mul [n0,n1,n2,...]",
	Short:     "multiplies numbers",
	Long: `
the mul subcommand multiplies together successive strings of numbers.
`,
	Run: runMul,
}

func runMul(cmd *Command, args []string) {
	if mulF {
		acc := 1.0
		for _, s := range args {
			v, err := strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
				os.Exit(3)
			}
			acc *= v
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
			acc *= v
		}
		fmt.Fprintln(os.Stdout, acc)
	}
}

func init() {
	cmdMul.Flag.BoolVar(&mulF, "f", false, "")
	addCommand(cmdMul)
}
