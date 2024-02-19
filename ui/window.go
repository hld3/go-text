package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateWindow() {
	a := app.New()
	w := a.NewWindow("Crappy text editor")
	w.Resize(fyne.Size{Width: 500, Height: 500})

	//TODO make a working menu some day.
	newItem := fyne.NewMenuItem("New", func() { })
	fileMenu := fyne.NewMenu("File", newItem)
	mainMenu := fyne.NewMainMenu(fileMenu)
	w.SetMainMenu(mainMenu)

	textEntry := widget.NewMultiLineEntry()
	textBox := container.NewStack(textEntry)
	w.SetContent(textBox)
	w.ShowAndRun()
}
