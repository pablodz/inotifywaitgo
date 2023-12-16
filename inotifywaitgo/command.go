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
		"inotifywait",
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
			if !Contains(VALID_EVENTS, int(event)) {
				return nil, errors.New(INVALID_EVENT)
			}
			baseCmd = append(baseCmd, EVENT_MAP[int(event)]+" ")
		}
	}

	// remove spaces on all elements
	var outCmd []string
	for _, v := range baseCmd {
		outCmd = append(outCmd, strings.TrimSpace(v))
	}

	if s.Verbose {
		fmt.Println("baseCmd:", outCmd)
	}

	return outCmd, nil
}

// function that checks if a string is in a slice of strings
func Contains[T string | int](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
