package cli

import (
	"fmt"

	"github.com/ixismail/nukeport/internals/colors"
)

func printHelp() {
	fmt.Println(" ")
	fmt.Printf("%sNukeport%s - Fast and cross-platform port killer\n\n", colors.Cyan, colors.Reset)
	fmt.Printf("%sUSAGE:%s\n", colors.Yellow, colors.Reset)
	fmt.Printf("  nuke port <port>\n\n")
	fmt.Printf("%sFLAGS:%s\n", colors.Yellow, colors.Reset)
	fmt.Printf("  -h, --help     Show help information\n")
	fmt.Printf("  -v, --version  Show current version\n\n")
	fmt.Printf("%sEXAMPLE:%s\n", colors.Yellow, colors.Reset)
	fmt.Printf("  nuke port 8080\n")
	fmt.Println(" ")
}
