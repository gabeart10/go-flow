package main

import (
	"github.com/nsf/termbox-go"
	"math"
)

type screen struct {
	width  int
	height int
	boxes  textBoxes
}

func newScreen() *screen {
	w, h := termbox.Size()
	return &screen{
		width:  w,
		height: h,
		boxes:  make(textBoxes, math.Ceil(w/3*h/3)),
	}
}

func (s *screen) addTextBox(box *textBox) {
	s.boxes = append(s.boxes, box)
}