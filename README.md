# OCR Text Scanner

Text OCR Scanner Software Tool. It is for extracting text from images and PDF files.

This software is __UNDER HEAVY DEVELOPMENT__. Use [this OCR software](https://abanoubhanna.com/posts/ocr) for now.

v0.1.0<sup>[1](semanticVersioning.md)</sup>

## commands

```sh
# get all dependencies of the project
go mod tidy

# build: produce an executable file/CLI app named "ocr"
go build -o ocr .

# run the app to identify a sample image
./ocr --lang=eng --img=img/default.png

# deps, build, run
go mod tidy && go build -o ocr . && ./ocr --lang=eng --img=img/default.png

# detect race conditions & memory leaks
go run -race .
```

## Roadmap & Tasks

Done & TO-DO of features, goals and values. tiny steps towards the goal.

- v0.1.0
  - OCR library used is [tesseract-OCR](https://github.com/tesseract-ocr/tesseract) with the [gosseract](https://github.com/otiai10/gosseract) v2 wrapper.
  - image to text
  - support English OCR
  - support Arabic OCR
  - add tests
  - use some test images from [renard314/textfairy](https://github.com/renard314/textfairy)
- Next
  - PDF to text
  - PDF to docx
  - PDF to selectable-text PDF
  - __scalable__ : take advantage of all CPU cores to get the job done faster
  - __bulk__ / __patch-processing__ : coroutines and parallelism for tasks / jobs
  - composable CLI app for scripts and automation
  - UX / easy to use / user friendly
  - cut the image into pieces/segments and concurrently OCR them. (performance)
  - available on Debian & Debian-based distros
  - available as snap
  - available as flatpak
  - available on Elementary OS
  - available on Mac OS (likely via HomeBrew)
  - available on Windows
  - scan history : { original image, processed image, extracted text }
  - migrate from opencv into Go libs

### References

- [3b1b : what is a convolution?](https://youtu.be/KuXjwB4LzSA)
- [train / refine tesseract OCR](https://github.com/abanoubha/train-tesseract-ocr)
- [tesseractfonts : fine tune tesseract for new fonts](https://github.com/dhivehi/tesseractfonts)
- [Pre-Recognize Library](https://github.com/leha-bot/PRLib) - library with algorithms for improving OCR quality.

## source code

The source code of OCR project can be found on:

- GitHub: <https://github.com/abanoubha/ocr.git>
- GitLab: <https://gitlab.com/abanoubha/ocr.git>
- CodeBerg: <https://codeberg.org/abanoubha/ocr.git>
