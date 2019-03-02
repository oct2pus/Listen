package trackbox

import (
	"listen/gui"
)

// SetValues handles all widget values that belong to GUI.TrackBox.
func SetValues(w gui.GUI) gui.GUI {
	// TrackBox
	w.TrackBox.Add(w.ImgTrack)
	w.TrackBox.ShowAll()

	return w
}

