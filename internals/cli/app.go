package cli

import (
	"fmt"
	"os"

	"github.com/ixismail/nukeport/internals/colors"
	"github.com/ixismail/nukeport/internals/nuker"
)

const Version = "1.0.0"

func Run(args []string) {

	if len(args) >= 2 && (args[1] == "-v" || args[1] == "--version") {
		fmt.Printf("%sNukeport%s version %s\n", colors.Cyan, colors.Reset, Version)
		os.Exit(0)
	}

	if len(args) < 3 || args[1] != "port" {
		fmt.Printf("%s[ERROR]%s Invalid command.\n", colors.Red, colors.Reset)
		fmt.Printf("%s[USAGE]%s nuke port <port>\n", colors.Cyan, colors.Reset)
		fmt.Printf("%s[EXAMPLE]%s nuke port 8080\n", colors.Green, colors.Reset)
		os.Exit(1)
	}

	port := args[2]
	
	fmt.Printf("%s[*]%s Target locked: Port %s\n", colors.Cyan, colors.Reset, port)

	nuker.Nuke(port)
}