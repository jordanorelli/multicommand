package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

var commands = []*Command{}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	if args[0] == "help" {
		help(args[1:])
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Run != nil {
			cmd.Flag.Usage = func() { cmd.Usage() }
			if cmd.CustomFlags {
				args = args[1:]
			} else {
				cmd.Flag.Parse(args[1:])
				args = cmd.Flag.Args()
			}
			cmd.Run(cmd, args)
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown command %#q\n\n", args[0])
	usage()
}

func usage() {
	printUsage(os.Stderr)
	os.Exit(2)
}

// help implements the 'help' command
func help(args []string) {
	if len(args) == 0 {
		printUsage(os.Stdout)
		// not exit 2: succeeded at 'go help'.
		return
	}
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: multicommand help command\n\nToo many arguments given.\n")
		os.Exit(2) // failed at 'go help'
	}

	name := args[0]
	for _, cmd := range commands {
		if cmd.Name() == name {
			tmpl(os.Stdout, helpTemplate, cmd)
			// not exit 2: succeeded at 'go help cmd'.
			return
		}
	}
}

func printUsage(w io.Writer) {
	tmpl(w, usageTemplate, commands)
}

// addCommand adds a given Command object to our global commands object, which
// is a special []*Command
func addCommand(cmd *Command) {
	commands = append(commands, cmd)
}

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace})
	// t.Funcs(template.FuncMap{"trim": strings.TrimSpace, "capitalize": capitalize})
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

var usageTemplate = `This is a test command.
Usage:

    multicommand [commandname] [arg1] [arg2] [...]

Available commands:
{{range .}}{{if .Runnable}}
    {{.Name}}: {{.Short}}{{end}}{{end}}
`

var helpTemplate = `{{if .Runnable}}usage: go {{.UsageLine}}

{{end}}{{.Long | trim}}
`
