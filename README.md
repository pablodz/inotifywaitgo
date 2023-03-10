# inotifywaitgo

Binding for inotifywait in golang, Fetch any directory event in your linux server easily. Fsnotify alternative

- Works with mounted volumes in Docker linux containers

Author: pablodz


## Example

```go
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
```