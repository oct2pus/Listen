package widgets

import (
	"listen/audio"
	"listen/gui"
	"listen/gui/widgets/appwindow"
	"listen/gui/widgets/header"
	"listen/gui/widgets/popmenu"
	"listen/gui/widgets/trackbox"
	"listen/util"

	"github.com/gotk3/gotk3/gtk"
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

// PlayPressed is what happens when the play button is pressed
func PlayPressed(g gui.GUI) {
	iconName, _ := g.ImgPlay.GetIconName()
	switch iconName {
	case "media-playback-start-symbolic":
		mus := audio.Read()

		g.ImgTrack.SetFromPixbuf(mus.Art)
		g.ImgPlay.SetFromIconName("media-playback-stop-symbolic",
			gtk.ICON_SIZE_BUTTON)
	case "media-playback-stop-symbolic":
		//audio.Stop()
		g.ImgPlay.SetFromIconName("media-playback-start-symbolic",
			gtk.ICON_SIZE_BUTTON)
	}
}

// FilePressed is what happens when a file button is pressed, returns
// a string with an absolute path to the file in it.
// Because everything seems to want that  and not an os.File because
// fuck byte slices I guess.
func FilePressed(g gui.GUI) {
	// create temporary holder window
	win, err := gtk.WindowNew(gtk.WINDOW_POPUP)
	if err != nil {
		util.SendError(err, "holder window")
	}
	defer win.Close()

	diag, err := gtk.FileChooserNativeDialogNew("Open File",
		win,
		gtk.FILE_CHOOSER_ACTION_OPEN,
		"Open",
		"Cancel")
	if err != nil {
		util.SendError(err, "file dialog")
	}

	if diag.Run() == int(gtk.RESPONSE_ACCEPT) {
		gui.MusicFile = diag.GetFilename()
		mus := audio.Read()

		g.ImgTrack.SetFromPixbuf(mus.Art)
	}
}
