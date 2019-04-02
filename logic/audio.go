package logic

import (
	"errors"
	"os"
	"time"

	"github.com/dhowden/tag"
	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
)

// Read reads and creates an Audio file
func Read(path string) AudioData {
	// Open file
	println("here")
	f, err := os.Open(path)
	if err != nil {
		SendError(err, "cannot open file")
	}
	var a AudioData
	var ssc beep.StreamSeekCloser
	var format beep.Format
	a.Path = path
	switch findFileType(a) {
	case tag.FLAC:
		ssc, format, err = flac.Decode(f)
	case tag.MP3:
		ssc, format, err = mp3.Decode(f)
	case tag.OGG:
		ssc, format, err = vorbis.Decode(f)
	default:
		SendError(errors.New("Invalid stream type"), "cannot read file")
		return AudioData{}
	}

	a = NewAudioData(&ssc, path)
	// Play Audio
	speaker.Init(format.SampleRate, format.SampleRate.N((time.Second / 10)))
	speaker.Play(a.Vol)
	return a
}

func findFileType(a AudioData) tag.FileType {
	f, meta := a.openMusic()
	defer f.Close() // better safe than sorry
	return meta.FileType()
}
