package main

import (
	"fmt"
	"log"

	"github.com/pablodz/inotifywaitgo/inotifywaitgo"
)

func main() {
	dir := "./"
	events := make(chan inotifywaitgo.FileEvent)
	errors := make(chan error)

	go inotifywaitgo.WatchPath(&inotifywaitgo.Settings{
		Dir:        dir,
		FileEvents: events,
		ErrorChan:  errors,
		Options: &inotifywaitgo.Options{
			Recursive: true,
			Events: []inotifywaitgo.EVENT{
				inotifywaitgo.CLOSE_WRITE,
			},
			Monitor: true,
		},
		Verbose: true,
	})

loopFiles:
	for {
		select {
		case event := <-events:
			// For each file close_write event usually there are 2 events,
			// is recommended to test inotifywait first
			log.Printf("[Event]%s, %v\n", event.Filename, event.Events)

			for _, e := range event.Events {
				switch e {
				case inotifywaitgo.CLOSE_WRITE:
					fmt.Printf("File %s close_write\n", event.Filename)
				case inotifywaitgo.CLOSE:
					fmt.Printf("File %s closed\n", event.Filename)
				}
			}

		case err := <-errors:
			fmt.Printf("Error: %s\n", err)
			break loopFiles
		}
	}
}
