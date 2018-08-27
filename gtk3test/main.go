package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	const appID = "org.gtk.example"
	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal("Could not create application.", err)
	}

	app.Connect("activate", func() {
		win, err := gtk.ApplicationWindowNew(app)
		if err != nil {
			log.Fatal("Could not create application window.", err)
		}

		win.SetTitle("Basic Application.")
		win.SetDefaultSize(400, 400)
		win.Show()
	})

	app.Run(os.Args)
}
