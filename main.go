// Package main launches the calculator app
package main

import "fyne.io/fyne/v2/app"

func main() {
	app := app.New()
	//app.SetIcon(resourceIconPng)

	c := gmeter{}
	c.loadUI(app)
	app.Run()
}