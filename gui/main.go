package gui

import (
	"image"
	"os"

	"github.com/disintegration/imaging"
	gosseract "github.com/otiai10/gosseract/v2"
)

// ocr-gui
func main() {

	// extracted, err := ocr(*img, *lang, false)
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
