# OCR Text Scanner

Text OCR Scanner Software Tool. It is for extracting text from images and PDF files.

v0.1.0<sup>[1](semanticVersioning.md)</sup>

![first iteration](./screenshots/1st-iteration.png)

## Done & TO-DO : features, goals and values

- [x] OCR library used is [tesseract-OCR](https://github.com/tesseract-ocr/tesseract) with the [gosseract](https://github.com/otiai10/gosseract) v2 wrapper.
- [x] image to text
- [ ] __scalable__ : take advantage of all CPU cores to get the job done faster
- [ ] __bulk__ / __patch-processing__ : coroutines and parallelism for tasks / jobs
- [ ] composable CLI app for scripts and automation
- [ ] UX / easy to use / user friendly
- [x] support English OCR
- [x] support Arabic OCR
- [x] add tests
- [x] add test images for binarization/thresholding
- [ ] cut the image into pieces/segments and concurrently OCR them. (performance)
- [ ] PDF to TXT
- [ ] PDF to docx
- [ ] PDF to selectable-text PDF
- [ ] available on Debian & Debian-based distros
- [ ] available as snap
- [ ] available as flatpak
- [ ] available on Elementary OS
- [ ] available on Mac OS (likely via HomeBrew)
- [ ] available on Windows

### References

- [3b1b : what is a convolution?](https://youtu.be/KuXjwB4LzSA)
- [got test images from renard314/textfairy](https://github.com/renard314/textfairy)
- [train / refine tesseract OCR](https://github.com/abanoub-hanna/train-tesseract-ocr)
- [tesseractfonts : fine tune tesseract for new fonts](https://github.com/dhivehi/tesseractfonts)
