package main

import (
	"github.com/johnnylee/gtk"
)

func cb(w gtk.Widget) {
	println("Got a callback!")
}

func main() {
	gtk.Init()
	gtk.AddFromFile("test2.glade")

	w := gtk.GetWidget("window")
	w.SignalConnect("destroy", func(w gtk.Widget) { gtk.MainQuit() })
	w.Show()

	btn := gtk.GetWidget("helloButton")
	btn.SignalConnect("clicked", cb)
	btn.SetTooltipMarkup("<b>It's a tip!</b>")

	gtk.Main()
}
