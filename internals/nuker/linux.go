package nuker

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/ixismail/nukeport/internals/colors"
)

func NukeLinux(port string) {

	// Search for the port using Linux's native 'fuser' command
	// format it as "<port>/tcp".
	cmd := exec.Command("fuser", fmt.Sprintf("%s/tcp", port))

	// Used Output() instead of CombinedOutput() because fuser prints the
	// port name to stderr, but outputs the actual clean PID to stdout.
	output, err := cmd.Output()

	// Clean up the output to get just the numbers
	pid := strings.TrimSpace(string(output))

	if err != nil || pid == "" {
		fmt.Printf("%s[*]%s No listening process found on port %s.\n", colors.Cyan, colors.Reset, port)
		return
	}

	// fuser can sometimes return multiple PIDs if child processes are sharing the port.
	// We grab just the first one.
	pids := strings.Fields(pid)
	targetPid := pids[0]

	if !askForConfirmation(port, targetPid) {
		fmt.Printf("%s[*]%s Operation cancelled by user.\n", colors.Cyan, colors.Reset)
		return
	}

	killCmd := exec.Command("kill", "-9", targetPid)
	killErr := killCmd.Run()

	if killErr != nil {
		fmt.Printf("%s[ERROR]%s Failed to kill PID %s. You may need sudo permissions.\n", colors.Red, colors.Reset, targetPid)
		return
	}

	fmt.Printf("%s[*]%s BOOM! Process %s nuked successfully.\n", colors.Green, colors.Reset, targetPid)
}