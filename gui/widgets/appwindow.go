package widgets

import (
	"listen/gui"
	"listen/logic"
)

// setAppWindowValues initializes all widget values that belong to GUI.AppWindow.
func setAppWindowValues(w gui.GUI) gui.GUI {
	// AppWindow
	w.AppWindow.Add(w.TrackBox)
	w.AppWindow.SetTitlebar(w.Header)
	w.AppWindow.SetDefaultSize(logic.ArtSize, logic.ArtSize)
	w.AppWindow.SetResizable(false) // may consider changing this one day
	// 								   don't think this application is ever
	// 								   supposed to go full screen tho
	w.AppWindow.Show()
	return w
}
