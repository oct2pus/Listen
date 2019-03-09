package logic

import (
	//	"github.com/dhowden/tag"

	"errors"
	"os"

	"github.com/gotk3/gotk3/gdk"
)

// FindArt takes the art from the track and returns a pixbuf of it
func FindArt(a AudioData) (*gdk.Pixbuf, error) {

	// process picture
	mus, meta := a.openMusic()
	defer mus.Close()

	pic := meta.Picture()

	// pic doesn't return an error, it only nils
	if pic == nil {
		return nil, errors.New("No picture found")
	}

	f, err := os.Create("./.temp_cover")
	if err != nil {
		SendError(err, "creating file")
	}
	defer f.Close()
	defer os.Remove("./.temp_cover")

	_, err = f.Write((*pic).Data)
	if err != nil {
		SendError(err, "writing to cover")
	}

	pix, err := gdk.PixbufNewFromFileAtScale("./.temp_cover",
		ArtSize,
		ArtSize,
		true)
	if err != nil {
		SendError(err, "setting album art pixbuf")
	}

	return pix, nil
}
