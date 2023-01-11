package inotifywaitgo

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

// Function that starts watching a path for new files and returns the file name (abspath) when a new file is finished writing
func WatchPath(s *Settings) {
	// Check if inotifywait is installed
	ok, err := checkDependencies()
	if !ok || err != nil {
		s.ErrorChan <- []byte(NOT_INSTALLED)
		return
	}

	// check if dir exists
	_, err = os.Stat(s.Dir)
	if os.IsNotExist(err) {
		s.ErrorChan <- []byte(DIR_NOT_EXISTS)
		return
	}

	// Stop any existing inotifywait processes
	if s.KillOthers {
		killOthers()
	}

	// Generate bash command
	cmdString, err := GenerateBashCommands(s)
	if err != nil {
		s.ErrorChan <- []byte(err.Error())
		return
	}

	// Start inotifywait in the input directory and watch for close_write events
	cmd := exec.Command(cmdString[0], cmdString[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		s.ErrorChan <- []byte(err.Error())
		return
	}
	if err := cmd.Start(); err != nil {
		s.ErrorChan <- []byte(err.Error())
		return
	}

	// Read the output of inotifywait and split it into lines
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			s.ErrorChan <- []byte(INVALID_OUTPUT)
			continue
		}

		// Extract the input file name from the inotifywait output
		prefix := parts[0]
		file := parts[len(parts)-1]
		// Send the file name to the channel
		s.OutFiles <- []byte(prefix + file)
	}
}
