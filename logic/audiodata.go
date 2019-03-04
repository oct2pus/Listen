package logic

import (
	"fmt"
	"os"

	"github.com/dhowden/tag"
	"github.com/faiface/beep"
	"github.com/gotk3/gotk3/gdk"
)

// AudioData is a wrapper around everything we need to store about an audio file.
type AudioData struct {
	EndTime int
	NowTime int
	Art     *gdk.Pixbuf
	Stream  beep.StreamSeekCloser
	Path    string
}

// NewAudioData defines a new AudioData
func NewAudioData(ssc beep.StreamSeekCloser, path string) AudioData {
	var a AudioData
	a.NowTime = 0
	a.EndTime = ssc.Len()
	a.Stream = ssc
	a.Path = path
	a.Art = FindArt(a)
	return a
}

// String returns a string of AudioData, mostly for diagnostics
func (a AudioData) String() string {
	return fmt.Sprintf("Now: %x\nEnd: %x\nStream: %x\nArt: %x\nPath: %x",
		a.NowTime,
		a.EndTime,
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
