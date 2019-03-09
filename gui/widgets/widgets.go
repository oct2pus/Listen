package widgets

// main widgets file
// widgets has the initalizer and all the event calls
// maybe i should move those out of here?

import (
	"listen/gui"
)

// Define calls and initalizes the starting values of all widgets
// these are all left to their own seperate files in the widgets package
func Define(window gui.GUI) gui.GUI {
	window = setAppWindowValues(window)
	window = setHeaderValues(window)
	//	window = SetPopMenuValues(window)
	window = setTrackBoxValues(window)
	return window
}
