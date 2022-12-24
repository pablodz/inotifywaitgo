package main

import (
	"log"

	"github.com/pablodz/inotifywaitgo/inotifywaitgo"
)

func main() {

	dir := "./"
	files := make(chan []byte)
	errors := make(chan []byte)

	go inotifywaitgo.WatchPath(&inotifywaitgo.Settings{
		Dir:       dir,
		OutFiles:  files,
		ErrorChan: errors,
		Options: &inotifywaitgo.OptionsInotify{
			Recursive: true,
			Events:    []string{inotifywaitgo.EventCloseWrite},
			Monitor:   true,
		},
		Verbose: true,
	})

	log.Println("Watching for changes in", dir)
	log.Println("Press Ctrl+C to stop")

loopFiles:
	for {
		select {
		case file := <-files:
			println(string(file))
		case err := <-errors:
			println(string(err))
			break loopFiles
		}
	}
}
