package nuker

import (
	"fmt"
	"nukeport/internals/colors"
	"runtime"
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
