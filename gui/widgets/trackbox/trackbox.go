package trackbox

import (
	"listen/gui"

	"github.com/gotk3/gotk3/gdk"
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

func setSize(pix *gdk.Pixbuf) *gdk.Pixbuf {

	pix.ScaleSimple(400, 400, gdk.INTERP_BILINEAR)
	return pix
}
