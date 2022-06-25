package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/disintegration/imaging"
	gosseract "github.com/otiai10/gosseract/v2"
)

// ocr-gui
func main() {
	go func() {
		w := app.NewWindow()
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()

	// extracted, err := ocr(*img, *lang, false)
}

func run(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			title := material.H1(th, "Hello, Gio")
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon
			title.Alignment = text.Middle
			title.Layout(gtx)

			e.Frame(gtx.Ops)
		}
	}
}

func ocr(imgpath, lang string, isBlackBg bool) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	// client.Languages = []string{"eng", "ara"}
	client.SetLanguage(lang)

	if isBlackBg == true {
		imgIo, _ := os.Open(imgpath)
		imgDec, _, _ := image.Decode(imgIo)
		inverted := imaging.Invert(imgDec)
		imaging.Save(inverted, "./temp.jpg")
		defer os.Remove("./temp.jpg")
		client.SetImage("./temp.jpg")
	} else {
		client.SetImage(imgpath)
	}

	//boundingBox, _ := client.GetBoundingBoxes(PageIteratorLevel.RIL_SYMBOL)
	// boundingBox, err := client.GetBoundingBoxes(gosseract.RIL_SYMBOL)
	// if err != nil {
	// 	return "", err
	// }

	text, err := client.Text()
	// text, err := client.HOCRText()
	if err != nil {
		return "", err
	}
	return text, nil
}
