package nuker

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/ixismail/nukeport/internals/colors"
)

func NukeWindows(port string) {

	// Search for the port using the Windows 'netstat' command
	cmd := exec.Command("cmd", "/c", fmt.Sprintf("netstat -ano | findstr :%s", port))
	output, err := cmd.CombinedOutput()

	if err != nil || len(output) == 0 {
		fmt.Printf("%s[*]%s No active process found running on port %s.\n", colors.Cyan, colors.Reset, port)
		return
	}

	// Split up the output text to find the exact Process ID (PID)
	lines := strings.Split(string(output), "\n")
	var pid string

	for _, line := range lines {
		// Only want the process that is actively "LISTENING"
		if strings.Contains(line, "LISTENING") {
			// fields splits the line by spaces. The PID is always the last item on a Windows netstat line.
			fields := strings.Fields(line)
			if len(fields) >= 5 {
				pid = fields[len(fields)-1] 
				break
			}
		}
	}

	if pid == "" {
		fmt.Printf("%s[*]%s No listening process found on port %s.\n", colors.Cyan, colors.Reset, port)
		return
	}

	if !askForConfirmation(port, pid) {
		fmt.Printf("%s[*]%s Operation cancelled by user.\n", colors.Cyan, colors.Reset)
		return
	}

	// Kill the process using Windows 'taskkill'
	killCmd := exec.Command("taskkill", "/F", "/PID", pid)
	killErr := killCmd.Run()

	if killErr != nil {
		fmt.Printf("%s[ERROR]%s Failed to kill PID %s. You might need to run your terminal as Administrator.\n", colors.Red, colors.Reset, pid)
		return
	}

	fmt.Printf("%s[*]%s BOOM! Process %s nuked successfully.\n", colors.Green, colors.Reset, pid)
}