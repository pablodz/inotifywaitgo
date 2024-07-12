package inotifywaitgo

import (
	"bufio"
	"os/exec"
)

// CheckDependencies verifies if inotifywait is installed.
func checkDependencies() (bool, error) {
	cmd := exec.Command("bash", "-c", "which inotifywait")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return false, err
	}

	if err := cmd.Start(); err != nil {
		return false, err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}
