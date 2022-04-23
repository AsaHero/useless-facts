package main

import (
	"encoding/json"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Fact struct {
	Text string `json:text`
}

func GetUselessFact() (string, error) {
	client := new(http.Client)

	resp, err := client.Get("https://uselessfacts.jsph.pl/random.json?language=en")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	text, _ := ioutil.ReadAll(resp.Body)
	fact := new(Fact)
	json.Unmarshal(text, fact)
	return fact.Text, nil

}

func main() {
	a := app.New()
	w := a.NewWindow("Useless facts")
	w.Resize(fyne.NewSize(800, 600))

	text := canvas.NewText("Click the button to get useless fact :)", color.White)
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 24
	fact := widget.NewLabel("")
	fact.Wrapping = fyne.TextWrapWord
	button := widget.NewButton("GET", func() {
		uselessfact, err := GetUselessFact()
		if err != nil {
			dialog.ShowError(err,w)
		} else {
			fact.SetText(uselessfact)
		}
	})

	box := container.New(layout.NewCenterLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	vbox := container.New(layout.NewVBoxLayout(), text, box, fact)
	w.SetContent(vbox)
	w.ShowAndRun()
}
