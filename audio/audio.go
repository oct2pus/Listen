package audio

import (
	"listen/util"
	"listen/gui"
	"os"
	"time"
	"github.com/dhowden/tag"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/vorbis"
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
	f, err := os.Open(gui.MusicFile)
	if err != nil {
		util.SendError(err, "cannot open file")
	}
	var a beep.StreamSeekCloser
	var format beep.Format
	switch findFileType() {
	case tag.FLAC: 
		a, format, err = flac.Decode(f)
	case tag.MP3:
		a, format, err = mp3.Decode(f)
	case tag.OGG:
		a, format, err = vorbis.Decode(f)
	default:
		return Audio{}	// TODO: not this
	}
	if err != nil {
		util.SendError(err, "invalid file")
	}

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

func findFileType() tag.FileType{
	f, meta := openMusic()
	defer f.Close() // better safe than sorry
	return meta.FileType()
}

func openMusic() (*os.File, tag.Metadata) {
	mus, err := os.Open(gui.MusicFile)
	if err != nil {
		util.SendError(err, "reading file")
	}

	meta, err := tag.ReadFrom(mus)
	if err != nil {
		util.SendError(err, "writing file to metadata")
	}

	return mus, meta
}