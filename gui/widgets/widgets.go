package widgets

import (
	"listen/gui"
	"listen/gui/widgets/appwindow"
	"listen/gui/widgets/header"
	"listen/gui/widgets/popmenu"
	"listen/gui/widgets/trackbox"
)

// Define calls other functions that change widget properties
// I do it this way to help organize what is used where.
func Define(window gui.GUI) gui.GUI {
	window = popmenu.SetValues(window)
	window = trackbox.SetValues(window)
	window = header.SetValues(window)
	window = appwindow.SetValues(window)
	return window
}
