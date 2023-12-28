package mageextras

import (
	"os"
	"os/exec"
)

// DockerScout executes a docker security audit.
func DockerScout(args ...string) error {
	cmdName := "docker"

	cmdParameters := []string{cmdName}
	cmdParameters = append(cmdParameters, "scout")
	cmdParameters = append(cmdParameters, "cves")
	cmdParameters = append(cmdParameters, args...)

	cmd := exec.Command(cmdName)
	cmd.Args = cmdParameters
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
