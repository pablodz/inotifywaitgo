package inotifywaitgo

import (
	"errors"
	"fmt"
	"strings"
)

func GenerateBashCommands(s *Settings) ([]string, error) {

	if s.Options == nil {
		return nil, errors.New(OPT_NIL)
	}

	if s.Dir == "" {
		return nil, errors.New(DIR_EMPTY)
	}

	baseCmd := []string{
		"bash", "-c", "inotifywait",
	}

	if s.Options.Monitor {
		baseCmd = append(baseCmd, "-m")
	}

	if s.Options.Recursive {
		baseCmd = append(baseCmd, "-r")
	}

	baseCmd = append(baseCmd, s.Dir)

	if len(s.Options.Events) > 0 {
		baseCmd = append(baseCmd, "-e")
		for _, event := range s.Options.Events {
			// if event not in VALID_EVENTS
			if !contains(VALID_EVENTS, event) {
				return nil, errors.New(INVALID_EVENT)
			}
			baseCmd = append(baseCmd, event)
		}
	}
	if s.Verbose {
		fmt.Println("baseCmd:", baseCmd)
	}

	// join baseCmd from third to last element
	var outCmd []string
	outCmd = append(outCmd, baseCmd[0])
	outCmd = append(outCmd, baseCmd[1])
	outCmd = append(outCmd, strings.Join(baseCmd[2:], " "))

	return outCmd, nil

}

// function that checks if a string is in a slice of strings
func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
