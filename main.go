package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"go-stress-testing/model"
)


func main() {    
    
    myGmeter, _ := NewGmeter()
    response := model.Response{}
	myApp := app.New()
	myWin := myApp.NewWindow("Gmeter")
	myWin.Resize(fyne.NewSize(500, 200))
  
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
		
	content := container.NewVBox(urlBox, numberBox, concurrencyBox, testBtn, respBox)
	myWin.SetContent(content)
	myWin.ShowAndRun()
  }