package main

import (
	"fmt"

	"github.com/pablodz/inotifywaitgo/inotifywaitgo"
)

func main() {
	dir := "./"
	files := make(chan string)
	errors := make(chan error)

	go inotifywaitgo.WatchPath(&inotifywaitgo.Settings{
		Dir:       dir,
		OutFiles:  files,
		ErrorChan: errors,
		Options: &inotifywaitgo.Options{
			Recursive: true,
			Events:    []string{inotifywaitgo.EventCloseWrite},
			Monitor:   true,
		},
		Verbose: true,
	})

loopFiles:
	for {
		select {
		case file := <-files:
			fmt.Printf("File: %s\n", file)
		case err := <-errors:
			fmt.Printf("Error: %s\n", err)
			break loopFiles
		}
	}
}
