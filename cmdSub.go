package main

import (
	"fmt"
	"os"
	"strconv"
)

var subF bool // -f flag

var cmdSub = &Command{
	UsageLine: "sub [n0,n1,n2,....]",
	Short:     "subtracts numbers",
	Long: `
the sub subcommand subtracts numbers from left to right.
`,
	Run: runSub,
}

func runSub(cmd *Command, args []string) {
	if subF {
		acc := 0
		for _, s := range args {
			v, err := strconv.Atoi(s)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
				os.Exit(3)
			}
			acc -= v
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
	addCommand(cmdSub)
	cmdSub.Flag.BoolVar(&subF, "f", false, "subtract floating-point numbers intead of integers")
}
