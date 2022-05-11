package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"

	//"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/container"
)

func showGalleryApp(w fyne.Window) {
	//	a := app.New()
	//	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(800, 620))
	root_src := "Images"
	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}
	tabs := container.NewAppTabs()
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir()) //checking in folder directory or not
		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1] //spliting string after .  and 1th place eg abc.png so abc at 0th place and png at 1th place
			if extension == "png" || extension == "jpg" || extension == "jpeg" {
				image := canvas.NewImageFromFile(root_src + "\\" + file.Name())
				image.FillMode = canvas.ImageFillOriginal
				tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}
	} // at above code we are finding path of image
	/*
		w.SetContent(container.NewBorder(pannelContent, nil, nil, nil, tabs))
		w.Show()
	*/
	w.Resize(fyne.NewSize(500, 280))

	w.SetContent(
		container.NewBorder(DeskBtn, nil, nil, nil, tabs),
	)

	w.Show()

}
