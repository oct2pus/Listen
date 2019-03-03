package appwindow

import (
	"listen/gui"
)

// SetValues handles all widget values that belong to GUI.AppWindow.
func SetValues(w gui.GUI) gui.GUI {
	// AppWindow
	w.AppWindow.Add(w.TrackBox)
	w.AppWindow.SetTitlebar(w.Header)
	w.AppWindow.SetDefaultSize(gui.ArtSize, gui.ArtSize)
	w.AppWindow.Show()
	return w
}
