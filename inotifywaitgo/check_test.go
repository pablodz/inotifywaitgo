package inotifywaitgo

import (
	"testing"
)

func TestCheckDependencies(t *testing.T) {
	// Run the test only if inotifywait is installed
	hasInotifywait, err := checkDependencies()

	if err != nil {
		t.Errorf("Error checking dependencies: %v", err)
	}

	if !hasInotifywait {
		t.Skip("Skipping test as inotifywait is not installed.")
	}

}
