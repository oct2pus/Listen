package widgets

import (
	"listen/gui"

	"github.com/gotk3/gotk3/gtk"
)

// setTrackBoxValues initializes all widget values that belong to GUI.TrackBox.
func setTrackBoxValues(w gui.GUI) gui.GUI {
	// TrackBox
	w.TrackBox.Add(w.ImgTrack)
	w.ImgTrack.SetHExpand(true)
	w.ImgTrack.SetVExpand(true)
	w.ImgTrack.SetHAlign(gtk.ALIGN_CENTER)
	w.TrackBox.ShowAll()

	return w
}

