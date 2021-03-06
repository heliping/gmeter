package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"

	"go-stress-testing/model"
)

type gmeter struct {	
	window  fyne.Window
}

func (g *gmeter)loadUI(myApp fyne.App) {    
    
    myGmeter, _ := NewGmeter()
    response := model.Response{}
	
	myWin := myApp.NewWindow("Gmeter")
	myWin.Resize(fyne.NewSize(500, 200))
    


	toolbar := widget.NewToolbar(
	widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {}),	)

	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("input url")

	numberEntry := widget.NewEntry()
	numberEntry.SetPlaceHolder("input testNumber")

	concurrencyEntry := widget.NewEntry()
	concurrencyEntry.SetPlaceHolder("input concurrency")
    
	respEntry := widget.NewEntry()	
   
	urlBox := container.New(layout.NewFormLayout(), widget.NewLabel("url                 "), urlEntry)
	numberBox := container.New(layout.NewFormLayout(), widget.NewLabel("number        "), numberEntry)
	concurrencyBox := container.New(layout.NewFormLayout(), widget.NewLabel("concurrency  "), concurrencyEntry)
	respBox := container.New(layout.NewFormLayout(), widget.NewLabel("result"), respEntry)
    
	testBtn := widget.NewButton("Test", func() {		
		response = myGmeter.Stress(urlEntry.Text, concurrencyEntry.Text, numberEntry.Text)
		respEntry.SetText(fmt.Sprintf("QPS:%8.2f|AverageTime:%8.2f\n", 
		response.Qps, 
		response.AverageTime))

	})   
		
	content := container.NewVBox(toolbar,urlBox, numberBox, concurrencyBox, testBtn, respBox)
	myWin.SetContent(content)
	myWin.ShowAndRun()
  }