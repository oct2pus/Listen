package gui

import (
	"log"

	"github.com/oct2pus/listen/src/logic"

	"github.com/gotk3/gotk3/gtk"
)

// Elements is a collection of widgets used in the GUI
// Some widgets are commented out because they won't be needed until
// a later revision.
type Elements struct {
	AppWindow *gtk.ApplicationWindow
	Header    *gtk.HeaderBar
	TrackBox  *gtk.Box
	PlayButt  *gtk.Button
	VolButt   *gtk.VolumeButton
	FileButt  *gtk.Button
	//	MenuButt  *gtk.MenuButton
	ProgScale *gtk.Scale
	ImgPlay   *gtk.Image
	ImgFile   *gtk.Image
	//	ImgMenu   *gtk.Image
	ImgTrack *gtk.Image
	//	PopMenu   *gtk.Popover
	//	PopPick   *gtk.FileChooserButton

}

// New creates a Elements struct.
// Aka it creates all widgets needed for the application.
// **GO FREAKS OUT IF YOU MAKE A Elements POINTER.**
func (g Elements) New(app *gtk.Application) Elements {
	var err error
	g.AppWindow, err = gtk.ApplicationWindowNew(app)
	if err != nil {
		logic.SendError(err, "application window")
	}
	g.Header, err = gtk.HeaderBarNew()
	if err != nil {
		log.Fatal(err, "header bar")
	}
	g.TrackBox, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1)
	if err != nil {
		logic.SendError(err, "track box")
	}
	g.PlayButt, err = gtk.ButtonNew()
	if err != nil {
		logic.SendError(err, "play button")
	}
	g.ProgScale, err = gtk.ScaleNewWithRange(gtk.ORIENTATION_HORIZONTAL,
		0, 100, 0.2)
	if err != nil {
		logic.SendError(err, "progress bar")
	}
	g.VolButt, err = gtk.VolumeButtonNew()
	if err != nil {
		logic.SendError(err, "volume button")
	}
	/*	g.MenuButt, err = gtk.MenuButtonNew()
		if err != nil {
			logic.SendError(err, "menu button")
		} */
	g.ImgPlay, err = gtk.ImageNew()
	if err != nil {
		logic.SendError(err, "play button icon")
	}
	/*	g.ImgMenu, err = gtk.ImageNew()
		if err != nil {
			logic.SendError(err, "menu icon")
		} */
	g.ImgTrack, err = gtk.ImageNewFromIconName("action-unavailable-symbolic",
		gtk.ICON_SIZE_DND)
	if err != nil {
		logic.SendError(err, "Track Art")
	}
	g.ImgFile, err = gtk.ImageNew()
	if err != nil {
		logic.SendError(err, "file icon")
	}
	/*	g.PopMenu, err = gtk.PopoverNew(g.MenuButt)
		if err != nil {
			logic.SendError(err, "popover")
		} */
	/* g.PopPick, err = gtk.FileChooserButtonNew("Open File",
		gtk.FILE_CHOOSER_ACTION_OPEN)
	if err != nil {
		logic.SendError(err, "file chooser")
	} */
	g.FileButt, err = gtk.ButtonNew()
	if err != nil {
		logic.SendError(err, "file chooser button")
	}
	return g
}
