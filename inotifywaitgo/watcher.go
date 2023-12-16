package inotifywaitgo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Function that starts watching a path for new files and returns the file name (abspath) when a new file is finished writing
func WatchPath(s *Settings) {
	// Check if inotifywait is installed
	ok, err := checkDependencies()
	if !ok || err != nil {
		s.ErrorChan <- fmt.Errorf(NOT_INSTALLED)
		return
	}

	// check if dir exists
	_, err = os.Stat(s.Dir)
	if os.IsNotExist(err) {
		s.ErrorChan <- fmt.Errorf(DIR_NOT_EXISTS)
		return
	}

	// Stop any existing inotifywait processes
	if s.KillOthers {
		killOthers()
	}

	// Generate bash command
	cmdString, err := GenerateBashCommands(s)
	if err != nil {
		s.ErrorChan <- err
		return
	}

	// Start inotifywait in the input directory and watch for close_write events
	cmd := exec.Command(cmdString[0], cmdString[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		s.ErrorChan <- err
		return
	}
	if err := cmd.Start(); err != nil {
		s.ErrorChan <- err
		return
	}

	// Read the output of inotifywait and split it into lines
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		log.Println(scanner.Text())
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			s.ErrorChan <- fmt.Errorf(INVALID_OUTPUT)
			continue
		}

		// Extract the input file name from the inotifywait output
		prefix := parts[0]
		file := parts[len(parts)-1]
		// Send the file name to the channel
		s.OutFiles <- fmt.Sprintf("%s%s", prefix, file)
	}
}
