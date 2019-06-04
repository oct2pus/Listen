package logic

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/dhowden/tag"
	"github.com/gotk3/gotk3/gdk"
)

// FindArt takes the art from the track and returns a pixbuf of it.
func FindArt(a AudioData) (*gdk.Pixbuf, error) {
	// process picture
	var pic *tag.Picture
	var pix *gdk.Pixbuf

	mus, meta := a.openMusic()
	defer mus.Close()

	if meta != nil {
		pic = meta.Picture()
	} else {
		pic = nil
	}

	loader, err := gdk.PixbufLoaderNew()
	if err != nil {
		SendError(err, "creating PixbufLoader")
	}
	// if we don't have a picture, look for cover.ext
	if pic == nil {
		cf, err := readFromCoverImg(mus)
		if err != nil {
			SendError(err, "image loading")
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
		pix, err = loader.WriteAndReturnPixbuf(coverData)
		if err != nil {
			SendError(err, "writing pixbuf")
		}
	} else {
		pix, err = loader.WriteAndReturnPixbuf((*pic).Data)
		if err != nil {
			SendError(err, "writing pixbuf")
		}
		pix, err = pix.ScaleSimple(ArtSize, ArtSize, gdk.INTERP_BILINEAR)
		if err != nil {
			SendError(err, "scaling pixbuf")
		}
	}

	return pix, nil
}

func readFromCoverImg(mus *os.File) (*os.File, error) {
	// formats supported by gdk.Pixbuf
	imgFormats := []string{".png", ".jpeg", ".jpg", ".tiff", ".tga", ".gif"}
	dir, _ := path.Split(mus.Name())
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, errors.New("cannot read directory")
	}
	for _, file := range files {
		for _, format := range imgFormats {
			if strings.Contains(file.Name(), format) {
				return os.Open(dir + file.Name())
			}
		}
	}

	return nil, errors.New("no cover image found")
}
