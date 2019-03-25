package logic

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/gotk3/gotk3/gdk"
)

// FindArt takes the art from the track and returns a pixbuf of it.
func FindArt(a AudioData) (*gdk.Pixbuf, error) {
	// process picture
	mus, meta := a.openMusic()
	defer mus.Close()

	pic := meta.Picture()

	// todo (possibly maybe) wrap pixbuf_new_from_bytes for gotk3
	f, err := os.Create("./.temp_cover")
	if err != nil {
		SendError(err, "creating file")
	}
	defer f.Close()

	defer os.Remove("./.temp_cover")

	// if we don't have a picture, look for cover.ext
	if pic == nil {
		cf, err := readFromCoverImg(mus)
		if err != nil {
			SendError(err, "no cover image, ignoring")
			return nil, err
		}
		defer cf.Close()
		fileInfo, err := cf.Stat()
		if err != nil {
			SendError(err, "could not read cover fileInfo, ignoring")
			return nil, err
		}
		coverData := make([]byte, fileInfo.Size())
		_, err = cf.Read(coverData)
		if err != nil {
			SendError(err, "reading from cover.ext")
		}
		f.Write(coverData)
	} else {
		_, err = f.Write((*pic).Data)
		if err != nil {
			SendError(err, "writing to cover")
		}
	}

	// this is done because i can't load images directly into pixbuf
	// from a byte stream
	pix, err := gdk.PixbufNewFromFileAtScale("./.temp_cover",
		ArtSize,
		ArtSize,
		true)

	if err != nil {
		SendError(err, "setting album art pixbuf")
	}

	return pix, nil
}

func readFromCoverImg(mus *os.File) (*os.File, error) {
	// formats supported by gdk.Pixbuf
	imgFormats := []string{".png", ".jpeg", ".jpg", ".tiff", ".tga", ".gif"}
	dir, _ := path.Split(mus.Name())
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, errors.New("Woah baby")
	}
	for _, file := range files {
		for _, format := range imgFormats {
			if strings.Contains(file.Name(), format) {
				return os.Open(dir + file.Name())
			}
		}
	}

	return nil, errors.New("No cover image found")
}
