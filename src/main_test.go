package main

import (
	"testing"
)

func TestOcr(t *testing.T) {
	result, err := ocr("/Users/mbp/projects/ocr/img/gill-sans-red.jpg", "eng", false)
	if err != nil {
		t.Error("error occurred")
	}
	if result != "Gill Sans red" {
		t.Error("result != Gill Sans red")
	}

	result, err = ocr("/Users/mbp/projects/ocr/img/gill-sans-black.jpg", "eng", false)
	if err != nil {
		t.Error("error occurred")
	}
	if result != "Gill Sans black" {
		t.Error("result != Gill Sans black")
	}

	result, err = ocr("/Users/mbp/projects/ocr/img/gill-sans-white.jpg", "eng", false)
	if err != nil {
		t.Error("error occurred")
	}
	if result != "Gill Sans white" {
		t.Error("result != Gill Sans white")
	}

	// black background
	result, err = ocr("/Users/mbp/projects/ocr/img/inverted.jpg", "ara", true)
	if err != nil {
		t.Error("error occurred")
	}
	if result != "صباح الخير" {
		t.Error(result, " (ar-black-white.jpg) != صباح الخير")
	}

	result, err = ocr("/Users/mbp/projects/ocr/img/ar-red-white.jpg", "ara", false)
	if err != nil {
		t.Error("error occurred")
	}
	if result != "صباح الخير" {
		t.Error(result, " (ar-red-white.jpg) != صباح الخير")
	}

	result, err = ocr("/Users/mbp/projects/ocr/img/ar-white-black.jpg", "ara", false)
	if err != nil {
		t.Error("error occurred")
	}
	if result != "صباح الخير" {
		t.Error(result, " (ar-white-black.jpg) != صباح الخير")
	}
}
