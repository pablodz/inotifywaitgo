package example

import "github.com/pablodz/inotifywaitgo/inotifywaitgo"

func Example() {
	dir := "./safasfsas"
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
