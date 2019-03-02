package util

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
)

// SendError is a helper function to write a formatted error.
func SendError(err error, issue string) {
	log.Fatal("Could not create "+issue+".", err)
}

// GetPixBuf gets a PixBuf from a file; hopefully scaling it properly.
func GetPixBuf(width int, height int) *gdk.Pixbuf {
	pix, err := gdk.PixbufNewFromFileAtScale("./cover.png", width, height, true)
	if err != nil {
		SendError(err, "file from image")
	}
	return pix
}
