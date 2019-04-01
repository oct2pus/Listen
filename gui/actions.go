package gui

import (
	"listen/logic"

	"github.com/faiface/beep/speaker"

	"github.com/gotk3/gotk3/gtk"
)

// Actions represents a coordination of the gui layer and the logic layer.
// aka Actions just holds everything needed to do a signal lol.
type Actions struct {
	GUI   Elements
	Audio logic.AudioData
}

// PlayPressed is occurs when the play button is pressed
// Pauses and Unpauses media playback.
func (a Actions) PlayPressed() Actions {
	iconName, _ := a.GUI.ImgPlay.GetIconName()
	a.startStop()
	switch iconName {
	case "media-playback-start-symbolic":
		a.GUI.ImgPlay.SetFromIconName("media-playback-stop-symbolic",
			gtk.ICON_SIZE_BUTTON)
	case "media-playback-stop-symbolic":
		a.GUI.ImgPlay.SetFromIconName("media-playback-start-symbolic",
			gtk.ICON_SIZE_BUTTON)
	}

	return a
}

// FilePressed occurs when you press the File button. Saves and then initalizes,
// a music stream and sets the trackart.
func (a Actions) FilePressed() Actions {
	// create temporary holder window
	win, err := gtk.WindowNew(gtk.WINDOW_POPUP)
	if err != nil {
		logic.SendError(err, "holder window")
	}
	defer win.Close()

	// create dialog
	diag, err := gtk.FileChooserNativeDialogNew("Open File",
		win,
		gtk.FILE_CHOOSER_ACTION_OPEN,
		"Open",
		"Cancel")
	if err != nil {
		logic.SendError(err, "file dialog")
	}

	// filter choices
	filt, err := gtk.FileFilterNew()
	if err != nil {
		logic.SendError(err, "file filter")
	}
	filt.AddPattern("*.mp3")
	filt.AddPattern("*.flac")
	filt.AddPattern("*.ogg") // could possibly break if not a vorbis file.
	diag.SetFilter(filt)

	// if option is selected
	if diag.Run() == int(gtk.RESPONSE_ACCEPT) {
		a.Audio = logic.Read(diag.GetFilename())
		a.GUI.ImgPlay.SetFromIconName("media-playback-stop-symbolic",
			gtk.ICON_SIZE_BUTTON)
		a.GUI.PlayButt.SetSensitive(true)
		a.GUI.VolButt.SetSensitive(true)
		a = a.VolumeSlid()
		if a.Audio.Art != nil {
			a.GUI.ImgTrack.SetFromPixbuf(a.Audio.Art)
		} else {
			a.GUI.ImgTrack.SetFromIconName("action-unavailable-symbolic",
				gtk.ICON_SIZE_DND)
		}
	}

	return a
}

// VolumeSlid is called when the volume slider is moved.
func (a Actions) VolumeSlid() Actions {
	speaker.Lock()
	val := a.GUI.VolButt.GetValue()
	switch int(val * 10) {
	default:
		a.Audio.Vol.Silent = false
		a.Audio.Vol.Volume = 0 - (1-val)*10
	case 0:
		a.Audio.Vol.Silent = true
	}
	speaker.Unlock()

	return a
}

func (a Actions) startStop() {
	speaker.Lock()
	a.Audio.Ctrl.Paused = !a.Audio.Ctrl.Paused
	speaker.Unlock()
}

// ParseArgs parses commandline arguments to launch a song.
func ParseArgs(args []string) {
	// TODO
	// This should probably be moved as well.
}
