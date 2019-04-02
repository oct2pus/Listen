package gui

import (
	"listen/logic"

	"github.com/faiface/beep/speaker"
	"github.com/gotk3/gotk3/gtk"
)

var block = false

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
	case START:
		a.GUI.ImgPlay.SetFromIconName(STOP,
			gtk.ICON_SIZE_BUTTON)
		block = false
	case STOP:
		a.GUI.ImgPlay.SetFromIconName(START,
			gtk.ICON_SIZE_BUTTON)
		block = true
	}

	return a
}

// FilePressed occurs when you press the File button. Saves and then initalizes,
// a music stream and sets the trackart.
func (a Actions) FilePressed() Actions {
	// create temporary holder window
	println("emitted")
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
	//	filt.AddPattern("*.flac")	// cannot seek
	//	filt.AddPattern("*.ogg") // could possibly break if not a vorbis file.
	diag.SetFilter(filt)

	// if option is selected
	if diag.Run() == int(gtk.RESPONSE_ACCEPT) {
		a = a.setup(diag.GetFilename())
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

// MoveProg occurs when the user moves the slider, occurs after left click is
// released.
func (a Actions) MoveProg() Actions {
	speaker.Lock()
	if int(a.GUI.ProgScale.GetValue()) == (*a.Audio.Stream).Len() {
		a.GUI.ProgScale.SetValue(float64((*a.Audio.Stream).Len() - 1))
	}
	err := (*a.Audio.Stream).Seek(int(a.GUI.ProgScale.GetValue()))
	if err != nil {
		logic.SendError(err, "Audio Stream Seek")
	}
	println((*a.Audio.Stream).Position(), a.GUI.ProgScale.GetValue())
	speaker.Unlock()
	block = false
	return a
}

// DrawProg occurs every moment a song is playing and not being blocked.
func (a Actions) DrawProg() Actions {
	if a.Audio.Stream != nil && !block {
		a.GUI.ProgScale.SetValue(float64((*a.Audio.Stream).Position()))
		a.GUI.ProgScale.QueueDraw()
		a = a.isEnd()
	}
	return a
}

func (a Actions) isEnd() Actions {
	if a.Audio.Stream != nil &&
		a.GUI.ProgScale.GetValue() == float64((*a.Audio.Stream).Len()) {

		var err error
		speaker.Lock()
		path := a.Audio.Path
		a.Audio = logic.AudioData{}
		speaker.Unlock() // speaker must be unlocked to play a stream
		a.Audio, err = logic.Read(path)
		if err != nil {
			return a
		}
		a.GUI.ProgScale.SetValue(0)
		a.startStop()
		a.GUI.ImgPlay.SetFromIconName(START,
			gtk.ICON_SIZE_BUTTON)
	}
	return a
}

// Block sets the block variable to true.
func (a Actions) Block() {
	block = true
}

// LoadFromCMD occurs when a file is supplied before openning the program.
func (a Actions) LoadFromCMD(path string) Actions {
	a = a.setup(path)
	return a
}

func (a Actions) setup(path string) Actions {
	var err error
	a.Audio, err = logic.Read(path)
	if err != nil {
		return a
	}
	a.GUI.ImgPlay.SetFromIconName(STOP,
		gtk.ICON_SIZE_BUTTON)
	a.GUI.ProgScale.SetSensitive(true)
	a.GUI.PlayButt.SetSensitive(true)
	a.GUI.VolButt.SetSensitive(true)
	a.GUI.ProgScale.SetRange(0, float64((*a.Audio.Stream).Len()))
	a = a.VolumeSlid()
	if a.Audio.Art != nil {
		a.GUI.ImgTrack.SetFromPixbuf(a.Audio.Art)
	} else {
		a.GUI.ImgTrack.SetFromIconName("action-unavailable-symbolic",
			gtk.ICON_SIZE_DND)
	}
	return a
}

func (a Actions) startStop() {
	speaker.Lock()
	a.Audio.Ctrl.Paused = !a.Audio.Ctrl.Paused
	speaker.Unlock()
}
