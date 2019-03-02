package audio

import (
	"listen/util"
	"os"
	"time"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	// beep might be used
	_ "github.com/faiface/beep"
)

// Audio is a wrapper around an os.File; hopefully its an audio file!

// Read reads and creates an Audio file
func Read() {
	f, err := os.Open("./android52 - 52 days of autumn - 07 Blue.mp3")
	if err != nil {
		util.SendError(err, "cannot open file")
	}
	audio, format, err := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N((time.Second / 10)))
	speaker.Play(audio)
}
