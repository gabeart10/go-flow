package main

import (
	"errors"
	"github.com/nsf/termbox-go"
)

const (
	cursorColor termbox.Attribute = termbox.ColorWhite
)

var cursor = &textBox{
	width:    1,
	height:   1,
	x:        0,
	y:        0,
	shown:    false,
	allCords: make([]cords, 1),
}

func checkCursorColl(s *screen) *textBox {
	cursor.allCords[0] = [2]int{cursor.x, cursor.y}
	return s.checkIfColliding(cursor)
}

func placeCursorAtXY(x, y int, s *screen) error {
	w := s.width
	buffer := termbox.CellBuffer()
	prevX := cursor.x
	prevY := cursor.y
	cursor.x = x
	cursor.y = y
	if checkCursorColl(s) == borderBox {
		cursor.x = prevX
		cursor.y = prevY
		return errors.New("placeCursorAtXY: XY is in a border")
	}
	if cursor.shown == false {
		termbox.SetCell(prevX, prevY, buffer[(prevY*w)+prevX].Ch, buffer[(prevY*w)+prevX].Fg, termbox.ColorDefault)
	}
	termbox.SetCell(x, y, buffer[(y*w)+x].Ch, buffer[(y*w)+x].Fg, cursorColor)
	return nil
}

func hideCursor(s *screen) error {
	w := s.width
	if cursor.shown == false {
		return errors.New("hideCursor: Cursor is hiden")
	}
	buffer := termbox.CellBuffer()
	termbox.SetCell(cursor.x, cursor.y, buffer[(cursor.y*w)+cursor.x].Ch, buffer[(cursor.y*w)+cursor.x].Fg, termbox.ColorDefault)
	return nil
}
