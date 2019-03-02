package trackbox

import (
	"listen/gui"
)

// SetValues handles all widget values that belong to GUI.TrackBox.
func SetValues(w gui.GUI) gui.GUI {
	// ImgTrack
	w.ImgTrack.SetFromPixbuf(setSize(w.ImgTrack.GetPixbuf()))

	// TrackBox
	w.TrackBox.Add(w.ImgTrack)
	w.TrackBox.ShowAll()

	return w
}

