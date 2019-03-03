package audio

import (
	"listen/util"
	"os"
	"time"
	"github.com/dhowden/tag"
	"github.com/faiface/beep/mp3"
	"github.com/gotk3/gotk3/gdk"

	"github.com/faiface/beep/speaker"

	"github.com/faiface/beep"
)

// Audio is a wrapper around an os.File; hopefully its an audio file!
type Audio struct {
	EndTime int
	NowTime int
	Art     *gdk.Pixbuf
	Stream  beep.StreamSeekCloser
	Path    string
}

// Read reads and creates an Audio file
func Read() Audio {
	// Open file
	f, err := os.Open("./android52 - 52 days of autumn - 07 Blue.mp3")
	if err != nil {
		util.SendError(err, "cannot open file")
	}
	case findFileType() {

	}
	a, format, err := mp3.Decode(f)

	audio := newAudio(a, f)

	// Play Audio
	speaker.Init(format.SampleRate, format.SampleRate.N((time.Second / 10)))
	speaker.Play(a)

	return audio
}

func newAudio(a beep.StreamSeekCloser, f *os.File) Audio {
	var audio Audio
	audio.NowTime = 0
	audio.EndTime = a.Len()
	audio.Stream = a
	audio.Art = findArt()
	return audio
}

func findFileType() {

}