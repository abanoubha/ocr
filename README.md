# OCR Text Scanner

Text OCR Scanner Software Tool. It is for extracting text from images and PDF files.

v0.1.0<sup>[1](#1)</sup> WIP<sup>[2](#2)</sup>

![first iteration](./screenshots/1st-iteration.png)

## features, goals and values

- __scalable__ : take advantage of all CPU cores to get the job done faster
- __bulk__ / __patch-processing__ : coroutines and parallelism for tasks / jobs
- __easy to use__ : user friendly
- All OS : available on all operating systems
- _intuitive_ graphical user interface
- composable CLI app for scripts and automation

## Done & TO-DO

- [x] standing on the shoulders of [giants](#3)<sup>3</sup>
- [x] image to text
- [x] support English OCR
- [x] support Arabic OCR
- [x] add tests
- [x] add test images for binarization/thresholding
- [ ] Debian
- [ ] snap
- [ ] flatpak
- [ ] Elementary OS (if possible)
- [ ] Mac OS (via HomeBrew)
- [ ] Windows (if possible)
- [ ] cut the image into pieces/segments and concurrently OCR them. (performance)
- [ ] PDF to TXT
- [ ] PDF to docx
- [ ] PDF to selectable-text PDF

### References

- <a id="1">[1]</a> [semantic versioning](semanticVersioning.md)
- <a id="2">[2]</a> Work in Progress.
- <a id="3">[3]</a> OCR library used is [tesseract-OCR](https://github.com/tesseract-ocr/tesseract) with the [gosseract](https://github.com/otiai10/gosseract) v2 wrapper.
- [3b1b : what is a convolution?](https://youtu.be/KuXjwB4LzSA)
- [got test images from renard314/textfairy](https://github.com/renard314/textfairy)
- [train / refine tesseract OCR](https://github.com/abanoub-hanna/train-tesseract-ocr)
- [tesseractfonts : fine tune tesseract for new fonts](https://github.com/dhivehi/tesseractfonts)
