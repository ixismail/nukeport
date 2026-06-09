package nuker

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/ixismail/nukeport/internals/colors"
)

func NukeMac(port string) {
	
	// Search for the port using 'lsof'
	// -i tcp:<port> targets the specific port
	// -sTCP:LISTEN ensures we only grab the background service, not a random active connection
	// -t strips all formatting and returns ONLY the Process ID
	cmd := exec.Command("lsof", "-i", fmt.Sprintf("tcp:%s", port), "-sTCP:LISTEN", "-t")
	output, err := cmd.CombinedOutput()

	// Clean up any invisible newline characters from the terminal output
	pid := strings.TrimSpace(string(output))

	if err != nil || pid == "" {
		fmt.Printf("%s[*]%s No listening process found on port %s.\n", colors.Cyan, colors.Reset, port)
		return
	}

	if !askForConfirmation(port, pid) {
		fmt.Printf("%s[*]%s Operation cancelled by user.\n", colors.Cyan, colors.Reset)
		return
	}

	killCmd := exec.Command("kill", "-9", pid)
	killErr := killCmd.Run()

	if killErr != nil {
		fmt.Printf("%s[ERROR]%s Failed to kill PID %s. You may need sudo permissions.\n", colors.Red, colors.Reset, pid)
		return
	}

	fmt.Printf("%s[*]%s BOOM! Process %s nuked successfully.\n", colors.Green, colors.Reset, pid)
}