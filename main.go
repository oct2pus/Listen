package main

import (
	"listen/gui"
	"listen/logic"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	const appID = "moe.jade.listen"
	args := os.Args
	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)

	if err != nil {
		logic.SendError(err, "application")
	}
	if args != nil {
		//		actions.ParseArgs(args)
	}
	app.Connect("activate", func() { activateConnect(app) })
	app.Run(os.Args)
}

func activateConnect(app *gtk.Application) {

	//create widgets
	var window gui.Elements
	window = gui.Elements.New(window, app)

	//define widgets
	window = gui.InitWidgets(window)

	// define actions
	actions := gui.Actions{GUI: window}

	window.PlayButt.Connect("clicked",
		func() { actions = actions.PlayPressed() })
	window.FileButt.Connect("clicked",
		func() { actions = actions.FilePressed() })
	window.VolButt.Connect("value-changed",
		func() { actions = actions.VolumeSlid() })

	// ProgScale Events
	window.ProgScale.Connect("button-press-event",
		func() { actions.Block() })
	window.ProgScale.Connect("button-release-event",
		func() { actions = actions.MoveProg() })
	window.ProgScale.ConnectAfter("draw",
		func() { actions = actions.DrawProg() })
	window.ProgScale.Connect("value-changed",
		func() { actions = actions.IsEnd() })
}
