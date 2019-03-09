package main

import (
	"listen/gui"
	"listen/gui/actions"
	"listen/gui/widgets"
	"listen/logic"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	const appID = "moe.jade.listen"

	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)

	if err != nil {
		logic.SendError(err, "application")
	}
	app.Connect("activate", func() { activateConnect(app) })
	app.Run(os.Args)
}

func activateConnect(app *gtk.Application) {
	//create widgets
	var window gui.GUI
	window = gui.GUI.New(window, app)

	//define widgets
	window = widgets.Define(window)

	//	window.PlayButt.Connect("clicked", func() { widgets.PlayPressed(window) })
	window.FileButt.Connect("clicked",
		func() { actions.FilePressed(window) })
}
