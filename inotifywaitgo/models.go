package inotifywaitgo

type Settings struct {
	// Directory to watch
	Dir string
	// Channel to send the file name to
	OutFiles chan []byte
	// Channel to send errors to
	ErrorChan chan []byte
	// Options for inotifywait
	Options *OptionsInotify
	// Kill other inotifywait processes
	KillOthers bool
	// verbose
	Verbose bool
}

type OptionsInotify struct {
	// Watch the specified file or directory.  If this option is not specified, inotifywait will watch the current working directory.
	Events []string
	// Print the name of the file that triggered the event.
	Format string
	// Watch all subdirectories of any directories passed as arguments.  Watches will be set up recursively to an unlimited depth.  Symbolic links are not traversed.  Newly created subdirectories will also be watched.
	Recursive bool
	// Set a time format string as accepted by strftime(3) for use with the `%T' conversion in the --format option.
	TimeFmt string
	// Instead of exiting after receiving a single event, execute indefinitely.  The default behaviour is to exit after the first event occurs.
	Monitor bool
}

const (
	// A watched file or a file within a watched directory was read from.
	EventAccess = "access"
	//A watched file or a file within a watched directory was written to.
	EventModify = "modify"
	// The metadata of a watched file or a file within a watched directory was modified.  This includes timestamps, file permissions, extended attributes etc.
	EventAttrib = "attrib"
	// A  watched  file or a file within a watched directory was closed, after being opened in writable mode.  This does not necessarily imply the file was written to.
	EventCloseWrite = "close_write"
	// A watched file or a file within a watched directory was closed, after being opened in read-only mode.
	EventCloseNowrite = "close_nowrite"
	//  A watched file or a file within a watched directory was closed, regardless of how it was opened.  Note that this  is  actually  implemented simply by listening for both close_write and close_nowrite, hence all close events received will be output as one of these, not CLOSE.
	EventClose = "close"
	// A watched file or a file within a watched directory was opened.
	EventOpen = "open"
	// A watched file or a file within a watched directory was moved to the watched directory.
	EventMovedTo = "moved_to"
	// A watched file or a file within a watched directory was moved from the watched directory.
	EventMovedFrom = "moved_from"
	// A watched file or a file within a watched directory was moved to or from the watched directory.  This is equivalent to listening for both moved_from and moved_to.
	EventMove = "move"
	// A watched file or directory was moved. After this event, the file or directory is no longer being watched.
	EventMoveSelf = "move_self"
	//  A file or directory was created within a watched directory.
	EventCreate = "create"
	// A watched file or a file within a watched directory was deleted.
	EventDelete = "delete"
	// A watched file or directory was deleted.  After this event the file or directory is no longer being watched.  Note that this event can  occur even if it is not explicitly being listened for.
	EventDeleteSelf = "delete_self"
	// The  filesystem  on  which  a  watched  file or directory resides was unmounted.  After this event the file or directory is no longer being 	watched.  Note that this event can occur even if it is not explicitly being listened to.
	EventUnmount = "unmount"
)

var VALID_EVENTS = []string{
	EventAccess,
	EventModify,
	EventAttrib,
	EventCloseWrite,
	EventCloseNowrite,
	EventClose,
	EventOpen,
	EventMovedTo,
	EventMovedFrom,
	EventMove,
	EventMoveSelf,
	EventCreate,
	EventDelete,
	EventDeleteSelf,
	EventUnmount,
}

/* ERRORS */
const NOT_INSTALLED = "inotifywait is not installed"
const OPT_NIL = "optionsInotify is nil"
const DIR_EMPTY = "directory is empty"
const INVALID_EVENT = "invalid event"
const INVALID_OUTPUT = "invalid output"
const DIR_NOT_EXISTS = "directory does not exists"
