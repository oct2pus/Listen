package actions

import (
	"listen/gui"
	"listen/logic"

	"github.com/gotk3/gotk3/gtk"
)

// PlayPressed is occurs when the play button is pressed
// Pauses and Unpauses media playback.
func PlayPressed(g gui.GUI) {
	iconName, _ := g.ImgPlay.GetIconName()
	switch iconName {
	case "media-playback-start-symbolic":
		//		logic.Continue()
		g.ImgPlay.SetFromIconName("media-playback-stop-symbolic",
			gtk.ICON_SIZE_BUTTON)
	case "media-playback-stop-symbolic":
		//		logic.Stop()
		g.ImgPlay.SetFromIconName("media-playback-start-symbolic",
			gtk.ICON_SIZE_BUTTON)
	}
}

// FilePressed is what happens when a file button is pressed, returns
// a string with an absolute path to the file in it.
// Because everything seems to want that and not an os.File because
// fuck byte slices I guess.
func FilePressed(g gui.GUI) {
	// create temporary holder window
	win, err := gtk.WindowNew(gtk.WINDOW_POPUP)
	if err != nil {
		logic.SendError(err, "holder window")
	}
	defer win.Close()

	diag, err := gtk.FileChooserNativeDialogNew("Open File",
		win,
		gtk.FILE_CHOOSER_ACTION_OPEN,
		"Open",
		"Cancel")
	if err != nil {
		logic.SendError(err, "file dialog")
	}
	filt, err := gtk.FileFilterNew()
	if err != nil {
		logic.SendError(err, "file filter")
	}
	filt.AddPattern("*.mp3")
	filt.AddPattern("*.flac")
	//	filt.AddPattern("*.ogg") could possibly break if its not a vorbis
	diag.SetFilter(filt)

	if diag.Run() == int(gtk.RESPONSE_ACCEPT) {
		mus := logic.Read(diag.GetFilename())
		g.ImgPlay.SetFromIconName("media-playback-stop-symbolic",
			gtk.ICON_SIZE_BUTTON)
		if mus.Art != nil {
			g.ImgTrack.SetFromPixbuf(mus.Art)
		} else {
			g.ImgTrack.SetFromIconName("action-unavailable-symbolic",
				gtk.ICON_SIZE_DND)
		}
	}
}
