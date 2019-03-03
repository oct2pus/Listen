package trackbox

import (
	"github.com/gotk3/gotk3/gtk"
	"listen/gui"
)

// SetValues handles all widget values that belong to GUI.TrackBox.
func SetValues(w gui.GUI) gui.GUI {
	// TrackBox
	w.TrackBox.Add(w.ImgTrack)
	w.ImgTrack.SetHExpand(true)
	w.ImgTrack.SetVExpand(true)
	w.ImgTrack.SetHAlign(gtk.ALIGN_CENTER)
	w.TrackBox.ShowAll()

	return w
}

