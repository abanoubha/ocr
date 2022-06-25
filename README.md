# Image to Text OCR Software

v0.1.0 pre-alpha (WIP) ([semantic versioning](semanticVersioning.md))

The OCR library used is [tesseract-OCR](https://github.com/tesseract-ocr/tesseract) with the [gosseract](https://github.com/otiai10/gosseract) v2 wrapper.

![first iteration](./screenshots/1st-iteration.png)

## features, goals and values

- scalable : take advantage of all CPU cores to get the job done faster
- bulk : coroutines and parallelism for tasks / jobs
- easy to use
- available on all operating systems
- intuitive graphical user interface
- composable CLI app for scripts and automation

## Done & TO-DO

- [x] image to text
- [x] support English OCR
- [x] support Arabic OCR
- [x] added tests
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
