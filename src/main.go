package main

import (
	"listen/src/gui"
	"listen/src/logic"
	"os"
	"strings"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	const appID = "moe.jade.listen"
	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_HANDLES_OPEN)

	if err != nil {
		logic.SendError(err, "application")
	}
	app.Connect("open", func() { attemptOpen(app) })
	app.Connect("activate", func() { activateConnect(app, false) })
	app.Run(os.Args)
}

func activateConnect(app *gtk.Application, cmdLaunch bool) {

	//create widgets
	var window gui.Elements
	window = gui.Elements.New(window, app)

	//define widgets
	window = gui.InitWidgets(window, cmdLaunch)

	// define actions
	actions := gui.Actions{GUI: window}
	if cmdLaunch {
		actions = actions.LoadFromCMD(os.Args[1])
	}

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
}

func attemptOpen(app *gtk.Application) {
	isMP3 := false
	if len(os.Args) > 1 {
		_, err := os.Stat(os.Args[1])
		if err == nil {
			if strings.HasSuffix(strings.ToLower(os.Args[1]), ".mp3") {
				isMP3 = true
			} else {
				logic.SendError(err, "not a mp3 file")
			}
		} else {
			logic.SendError(err, "invalid file, ignoring")
		}
	}
	activateConnect(app, isMP3)
}
