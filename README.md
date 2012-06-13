Multicommand is a reference command that basically takes the command-suite
aspect of the Go command and extracts it.  It's shared here for posterity.
(the full source to the Go command is stored in `$GOROOT/src/cmd/go`)

Four subcommands are naively implemented, representing a four-function
calculator with a `-f` switch for each function to toggle floating point
arithmetic.
