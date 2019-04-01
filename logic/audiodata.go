package logic

import (
	"fmt"
	"os"

	"github.com/dhowden/tag"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/gotk3/gotk3/gdk"
)

// AudioData is a wrapper around everything we need to store about an audio
// file.
type AudioData struct {
	Art    *gdk.Pixbuf
	Stream *beep.StreamSeekCloser
	Ctrl   *beep.Ctrl
	Vol    *effects.Volume
	Path   string
}

// NewAudioData defines a new AudioData.
func NewAudioData(ssc *beep.StreamSeekCloser, path string) AudioData {
	var a AudioData
	var err error
	a.Stream = ssc
	a.Ctrl = &beep.Ctrl{Streamer: *a.Stream, Paused: false}
	a.Vol = &effects.Volume{
		Streamer: a.Ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
	a.Path = path
	a.Art, err = FindArt(a)
	if err != nil {
		a.Art = nil
	}
	return a
}

// String returns a string of AudioData, this is for diagnostics.
func (a AudioData) String() string {
	return fmt.Sprintf("\nNow: %v\nEnd: %v\nStream: %v\nArt: %v\nPath: %v",
		(*a.Stream).Position(),
		(*a.Stream).Len(),
		a.Stream,
		a.Art,
		a.Path)
}

func (a AudioData) openMusic() (*os.File, tag.Metadata) {
	mus, err := os.Open(a.Path)
	if err != nil {
		SendError(err, "reading file")
	}

	meta, err := tag.ReadFrom(mus)
	if err != nil {
		SendError(err, "writing file to metadata")
	}

	return mus, meta
}
