package main

import (
	"flag"
	"fmt"
	"image"
	"os"
	"strconv"

	"github.com/disintegration/imaging"
	gosseract "github.com/otiai10/gosseract/v2"
	"gocv.io/x/gocv"
)

// ocr --lang=ara --img=xyz.png
func main() {

	// A/B Testing :
	// - use custom threshold vs the default threshold in Tesseract
	// - determine the better result by the overall confidence
	//
	// osf, _ := os.Open("../img/testar.jpg")
	// defer osf.Close()
	// omg, _ := jpeg.Decode(osf)
	// // result := OtsuThreshold(omg)
	// result := threshold(omg, (255/2)+50)
	// ff, _ := os.Create("result.jpg")
	// jpeg.Encode(ff, result, &jpeg.Options{Quality: 95})

	lang := flag.String("lang", "eng", " Language of the written text. eng or ara as a language.")
	img := flag.String("img", "", " Image.")
	flag.Parse()

	if *img == "" {
		fmt.Println("Usage : ocr --lang=eng --img=~/Downloads/xyz.png")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *lang != "eng" && *lang != "ara" {
		fmt.Println("--lang must be ara for Arabic or eng for English.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println(*lang, *img)
	extracted, err := ocr(*img, *lang, false)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

	if extracted == "" {
		fmt.Println("no text extracted! something went wrong")
	}

	fmt.Println(extracted)

	// thresholding
	adaptiveThresholdBin("0000", 61, 16)
	adaptiveThresholdOtsu("0000", 61, 16)

	adaptiveThresholdBin("0001", 41, 14)
	adaptiveThresholdOtsu("0001", 41, 14)

	adaptiveThresholdBin("0002", 61, 16)
	adaptiveThresholdOtsu("0002", 61, 16)

	adaptiveThresholdBin("0003", 11, 11)
	adaptiveThresholdOtsu("0003", 11, 11)

	adaptiveThresholdBin("0004", 11, 11)
	adaptiveThresholdOtsu("0004", 11, 11)

	adaptiveThresholdBin("0005", 31, 11)
	adaptiveThresholdOtsu("0005", 31, 11)

	adaptiveThresholdBin("0006", 31, 11)
	adaptiveThresholdOtsu("0006", 31, 11)
}

func ocr(imgpath, lang string, isBlackBg bool) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	// client.Languages = []string{"eng", "ara"}
	client.SetLanguage(lang)

	if isBlackBg {
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
} // ocr

// // implemented in Go
// func threshold(img image.Image, threshold float64) image.Image {
// 	// Get the image dimensions.
// 	width, height := img.Bounds().Dx(), img.Bounds().Dy()

// 	// Create a new image with the same dimensions as the original image.
// 	thresholded := image.NewGray(img.Bounds())

// 	// Set the pixels to black if they are below the threshold, and white otherwise.
// 	for y := 0; y < height; y++ {
// 		for x := 0; x < width; x++ {
// 			// Get the pixel value.
// 			// v := img.At(x, y).Y
// 			r, g, b, _ := img.At(x, y).RGBA()
// 			v := float64(r+g+b) / (3 * 255)

// 			// Set the pixel.
// 			if v < threshold {
// 				thresholded.Set(x, y, color.Black)
// 			} else {
// 				thresholded.Set(x, y, color.White)
// 			}
// 		}
// 	}

// 	// Return the thresholded image.
// 	return thresholded
// }

// func OtsuThreshold(img image.Image) image.Image {
// 	// Get the image dimensions.
// 	width, height := img.Bounds().Dx(), img.Bounds().Dy()

// 	// Create a new image with the same dimensions as the original image.
// 	thresholded := image.NewGray(img.Bounds())

// 	// Calculate the Otsu threshold.
// 	var total int64
// 	var sum float64
// 	var w0, w1 float64
// 	var u0, u1 float64
// 	var var0, var1 float64
// 	for y := 0; y < height; y++ {
// 		for x := 0; x < width; x++ {
// 			// Get the pixel value.
// 			// v := img.At(x, y).Y
// 			r, g, b, _ := img.At(x, y).RGBA()
// 			// v := float64(r+g+b) / (3 * 255)
// 			const rConst float64 = 0.2126
// 			const gConst float64 = 0.7152
// 			const bConst float64 = 0.0722
// 			v := (rConst*float64(r) + gConst*float64(g) + bConst*float64(b)) / 255

// 			// Update the total, sum, w0, w1, u0, u1, var0, and var1 variables.
// 			total += 1
// 			sum += v
// 			w0 += v < 0.5
// 			w1 += v >= 0.5
// 			u0 += v * (v < 0.5)
// 			u1 += v * (v >= 0.5)
// 			var0 += v * (v < 0.5) * (v < 0.5)
// 			var1 += v * (v >= 0.5) * (v >= 0.5)
// 		}
// 	}

// 	// Calculate the Otsu threshold.
// 	var threshold float64
// 	for t := 0.0; t <= 1.0; t += 0.01 {
// 		var m0, m1 float64
// 		m0 = u0 / w0
// 		m1 = u1 / w1
// 		var varBetween float64
// 		varBetween = w0 * w1 * (m0 - m1) * (m0 - m1) / total
// 		if varBetween > var0 {
// 			var0 = varBetween
// 			threshold = t
// 		}
// 	}

// 	// Set the pixels to black if they are below the threshold, and white otherwise.
// 	for y := 0; y < height; y++ {
// 		for x := 0; x < width; x++ {
// 			// Get the pixel value.
// 			v := img.At(x, y).Y

// 			// Set the pixel.
// 			if v < threshold {
// 				thresholded.At(x, y) = color.Black
// 			} else {
// 				thresholded.At(x, y) = color.White
// 			}
// 		}
// 	}

// 	// Return the thresholded image.
// 	return thresholded
// } // OtsuThreshold

// blockSize = 11,21,31,41,51,61,..
// c = 1,2,3,4,..
func adaptiveThresholdBin(imageName string, blockSize int, c float32) {
	inpath := "../img/" + imageName + ".png"
	img := gocv.IMRead(inpath, gocv.IMReadGrayScale)
	defer img.Close()

	// whitePixels := gocv.CountNonZero(img)
	// totalPixels := img.Rows() * img.Cols()
	// whitePercentage := float64(whitePixels) / float64(totalPixels) * 100.0

	threshold := gocv.NewMat()
	defer threshold.Close()

	gocv.AdaptiveThreshold(img, &threshold, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinary, blockSize, c)
	// if whitePercentage > 50.0 {
	// 	gocv.AdaptiveThreshold(img, &threshold, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinary, blockSize, c)
	// } else {
	// 	gocv.AdaptiveThreshold(img, &threshold, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinaryInv, blockSize, c)
	// }

	binary := gocv.NewMat()
	defer binary.Close()
	mat := gocv.NewMatFromScalar(
		gocv.Scalar{Val1: 127, Val2: 0, Val3: 0, Val4: 0},
		gocv.MatTypeCV16SC1,
	)
	gocv.Compare(threshold, mat, &binary, gocv.CompareGT)
	result := gocv.NewMat()
	defer result.Close()
	binary.ConvertTo(&result, gocv.MatTypeCV8UC1)
	outpath := "../img/" + imageName + "-adaptive-mean-bin-blk-" + strconv.Itoa(blockSize) + "-c-" + fmt.Sprintf("%0.0f", c) + ".png"
	gocv.IMWrite(outpath, result)
}

func adaptiveThresholdOtsu(imageName string, blockSize int, c float32) {
	inpath := "../img/" + imageName + ".png"
	img := gocv.IMRead(inpath, gocv.IMReadGrayScale)
	defer img.Close()

	whitePixels := gocv.CountNonZero(img)
	totalPixels := img.Rows() * img.Cols()
	whitePercentage := float64(whitePixels) / float64(totalPixels) * 100.0

	blur := gocv.NewMat()
	defer blur.Close()
	gocv.GaussianBlur(img, &blur, image.Point{5, 5}, 0, 0, gocv.BorderDefault)
	threshold := gocv.NewMat()
	defer threshold.Close()

	if whitePercentage > 50.0 {
		gocv.AdaptiveThreshold(img, &threshold, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinary, blockSize, c)
	} else {
		gocv.AdaptiveThreshold(img, &threshold, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinaryInv, blockSize, c)
	}

	binary := gocv.NewMat()
	defer binary.Close()
	mat := gocv.NewMatFromScalar(gocv.Scalar{Val1: 127, Val2: 0, Val3: 0, Val4: 0}, gocv.MatTypeCV16SC1)
	gocv.Compare(threshold, mat, &binary, gocv.CompareGT)
	result := gocv.NewMat()
	defer result.Close()
	binary.ConvertTo(&result, gocv.MatTypeCV8UC1)
	outpath := "../img/" + imageName + "-adaptive-mean-bin-blk-" + strconv.Itoa(blockSize) + "-c-" + fmt.Sprintf("%0.0f", c) + "-blur.png"
	gocv.IMWrite(outpath, result)
}
