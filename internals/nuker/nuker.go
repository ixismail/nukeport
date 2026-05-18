package nuker

import (
	"bufio"
	"fmt"
	"nukeport/internals/colors"
	"os"
	"runtime"
	"strings"
)


func Nuke(port string) {
	osType := runtime.GOOS

	switch osType {
	case "windows":
		fmt.Printf("%s[*]%s OS Detected: 💻 Windows\n", colors.Cyan, colors.Reset)
	case "darwin":
		fmt.Printf("%s[*]%s OS Detected: 🍎 macOS\n", colors.Cyan, colors.Reset)
	case "linux":
		fmt.Printf("%s[*]%s OS Detected: 🐧 Linux\n", colors.Cyan, colors.Reset)
	default:
		fmt.Printf("%s[ERROR]%s Unsupported OS -> %s\n", colors.Red, colors.Reset, osType)
	}

}

func askForConfirmation(port string, pid string) bool {
	
	fmt.Printf("%s[?]%s Found Process ID %s holding port %s. Kill it? [Y/n]: ", colors.Yellow, colors.Reset, pid, port)

	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = strings.ToLower(strings.TrimSpace(answer))

	if answer == "y" || answer == "yes" {
		return true
	}

	return false
}