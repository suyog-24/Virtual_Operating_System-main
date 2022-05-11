package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var count int = 1

func showTextEditor() {

	//	a := app.New()
	//	w := a.NewWindow("text editor")
	//	w.Resize(fyne.NewSize(600, 600))
	//
	w := myApp.NewWindow("TextE")
	//

	content := container.NewVBox( //vertual box container
		container.NewHBox( //horizontal box
			//Box
			//The box widget is a simple horizontal or vertical container that uses the box
			// layout to lay out the child components. You can pass the objects to include in the
			//container.NewHBox() or container.NewVBox() constructor functions.
			//It is also possible to add items to a box widget after it is created using Add()
			// (to after the existing content) or remove items using Remove().
			widget.NewLabel("New Text Editor"),
		),
	)
	content.Add(widget.NewButton("Add New File", func() {
		content.Add(widget.NewLabel("New File" + strconv.Itoa(count)))
		//strconv.Itoa ==this is use to convert int to char because we cant use int with func
		count++

	}))
	input := widget.NewMultiLineEntry()
	//widget.NewEntry()=for single line ,,NewMultiLineEntry()==multiline
	input.SetPlaceHolder("Enter text...")
	//input.Resize(fyne.NewSize(400, 400))

	//saving function
	saveBtn := widget.NewButton("Save file", func() {
		//funcnewfile save fuction
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text) //byte array for saving text file
				uc.Write(textData)
			}, w)
		//name to file
		saveFileDialog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")
		saveFileDialog.Show()

	})
	openBtn := widget.NewButton("Open Text File", func() {
		openFileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				ReadData, _ := ioutil.ReadAll(r)

				output := fyne.NewStaticResource("New File", ReadData)
				viewData := widget.NewMultiLineEntry()
				viewData.SetText(string(output.StaticContent))
				w := fyne.CurrentApp().NewWindow(
					string(output.StaticName))
				w.SetContent(container.NewScroll(viewData))
				w.Resize(fyne.NewSize(400, 400))
				w.Show()
			}, w)
		openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		openFileDialog.Show()
	})

	//w.SetContent(
	WContainer := container.NewVBox(
		container.NewVBox(
			content,
			input,

			container.NewHBox(
				saveBtn,
				openBtn,
			),
		),
	)
	//
	//w := myApp.NewWindow("Calculator")
	w.Resize(fyne.NewSize(500, 280))

	w.SetContent(
		container.NewBorder(DeskBtn, nil, nil, nil, WContainer),
	)

	w.Show()

	//	w.ShowAndRun()
}
