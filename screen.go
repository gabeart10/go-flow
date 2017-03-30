package main

import (
	"github.com/nsf/termbox-go"
)

const (
	modeResize mode = iota + 6
	modeEditBox
	modeMove
	modeMoveBox
)

type mode uint8

type screen struct {
	width      int
	height     int
	boxes      textBoxes
	screenMode mode
}

func newScreen() *screen {
	w, h := termbox.Size()
	h--
	tempScreen := &screen{
		width:      w,
		height:     h,
		boxes:      make(textBoxes, int(w/3*h/3)),
		screenMode: modeMove,
	}
	tempScreen.changeMode(modeMove)
	return tempScreen
}

func (s *screen) addTextBox(box *textBox) {
	s.boxes = append(s.boxes, box)
}

func (s *screen) changeMode(toMode mode) {
	for i := 0; i < s.width; i++ {
		termbox.SetCell(i, s.height, ' ', termbox.ColorDefault, termbox.ColorBlack)
	}
	text := " "
	switch toMode {
	case modeMove:
		text = "--move--"
		s.screenMode = modeMove
	case modeResize:
		text = "--resize--"
		s.screenMode = modeResize
	case modeMoveBox:
		text = "--movebox--"
		s.screenMode = modeMoveBox
	case modeEditBox:
		text = "--editbox--"
		s.screenMode = modeEditBox
	}
	i := 0
	for _, char := range text {
		termbox.SetCell(i, s.height, char, termbox.ColorBlue, termbox.ColorBlack)
		i++
	}
}
