package header

import (
	"listen/gui"

	"github.com/gotk3/gotk3/gtk"
)

// SetValues handles all widget values that belong to GUI.Header.
func SetValues(w gui.GUI) gui.GUI {
	// Buttons
	//w.ImgMenu.SetFromIconName("open-menu-symbolic", gtk.ICON_SIZE_BUTTON)
	w.ImgFile.SetFromIconName("document-open-symbolic", gtk.ICON_SIZE_BUTTON)
	w.ImgPlay.SetFromIconName("media-playback-start-symbolic",
		gtk.ICON_SIZE_BUTTON)
	//w.MenuButt.SetImage(w.ImgMenu)
	w.PlayButt.SetImage(w.ImgPlay)
	w.FileButt.SetImage(w.ImgFile)
	//w.MenuButt.SetPopover(w.PopMenu)

	// ProgScale
	w.ProgScale.SetDrawValue(false)
	w.ProgScale.SetHExpand(true)

	// Header
	// *Remember: Order Matters!*
	w.Header.SetCustomTitle(w.ProgScale)
	w.Header.PackStart(w.PlayButt)
	//w.Header.PackEnd(w.MenuButt)
	w.Header.PackEnd(w.FileButt)
	w.Header.PackEnd(w.VolButt)
	w.Header.SetShowCloseButton(true)
	w.Header.ShowAll()

	return w
}
