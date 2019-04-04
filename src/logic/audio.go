package logic

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/dhowden/tag"
	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
)

// Read reads and creates an Audio file
func Read(path string) (AudioData, error) {
	// Open file
	f, err := os.Open(path)
	if err != nil {
		SendError(err, "cannot open file")
		return AudioData{}, err
	}
	var a AudioData
	var ssc beep.StreamSeekCloser
	var format beep.Format
	a.Path = path
	musTag := findFileType(a)
	switch musTag {
	// TODO: FLAC needs to be able to seek.
	case tag.FLAC:
		ssc, format, err = flac.Decode(f)
	case tag.MP3:
		ssc, format, err = mp3.Decode(f)
	// TODO: I need to detect vorbis encoding.
	case tag.OGG:
		ssc, format, err = vorbis.Decode(f)
	default:
		err = errors.New("invalid stream type")
		SendError(err, "cannot read file")
		return AudioData{}, err
	}
	if err != nil {
		SendError(err, "cannot read file")
		return AudioData{}, err
	}

	a = NewAudioData(&ssc, path)
	// Play Audio
	speaker.Init(format.SampleRate, format.SampleRate.N((time.Second / 10)))
	speaker.Play(a.Vol)
	return a, nil
}

func findFileType(a AudioData) tag.FileType {
	f, meta := a.openMusic()
	defer f.Close() // better safe than sorry
	if meta == nil || meta.FileType() == tag.UnknownFileType {
		return noID3Search(f.Name())
	}
	return meta.FileType()
}

func noID3Search(s string) tag.FileType {
	exts := []string{".mp3", ".ogg", ".flac"}
	for _, ext := range exts {
		if strings.Contains(s, ext) {
			return tag.FileType(strings.ToUpper(strings.TrimPrefix(ext, ".")))
		}
	}
	return tag.FileType("")
}
