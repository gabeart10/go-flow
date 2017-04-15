package main

import (
	"github.com/nsf/termbox-go"
	"os"
)

var selectedTextBox *textBox = nil

func (s *screen) startEventWatcher() {
	for {
		event := termbox.PollEvent()
		switch s.screenMode {
		case modeMove:
			s.prossMove(event)
		case modeResize:
			s.prossResize(event)
		case modeMoveBox:
			s.prossMoveBox(event)
		case modeEditBox:
			//s.prossEditBox(event)
		}
	}
}

func exitProgram() {
	termbox.Close()
	os.Exit(3)
}

func (s *screen) prossMove(event termbox.Event) {
	if event.Ch == 0 && event.Key == termbox.KeyCtrlC {
		exitProgram()
	}
	switch event.Ch {
	case 'h':
		placeCursorAtXY(cursor.x-1, cursor.y, s)
	case 'k':
		placeCursorAtXY(cursor.x, cursor.y-1, s)
	case 'j':
		placeCursorAtXY(cursor.x, cursor.y+1, s)
	case 'l':
		placeCursorAtXY(cursor.x+1, cursor.y, s)
	case 'm':
		collData := checkCursorColl(s)
		if collData != nil {
			s.changeMode(modeMoveBox)
			selectedTextBox = collData
			hideCursor(s)
		}
	case 'r':
		collData := checkCursorColl(s)
		if collData != nil {
			s.changeMode(modeResize)
			selectedTextBox = collData
			hideCursor(s)
		}
	case 'a':
		s.createAndAddTextBox()
	}
	termbox.Flush()
}

func (s *screen) prossMoveBox(event termbox.Event) {
	switch event.Key {
	case termbox.KeyCtrlC:
		exitProgram()
	case termbox.KeyCtrlLsqBracket:
		s.changeMode(modeMove)
		placeCursorAtXY(selectedTextBox.x, selectedTextBox.y, s)
		selectedTextBox = nil
	}
	switch event.Ch {
	case 'h':
		selectedTextBox.hide()
		selectedTextBox.placeAtXY(selectedTextBox.x-1, selectedTextBox.y, s, true)
	case 'j':
		selectedTextBox.hide()
		selectedTextBox.placeAtXY(selectedTextBox.x, selectedTextBox.y+1, s, true)
	case 'k':
		selectedTextBox.hide()
		selectedTextBox.placeAtXY(selectedTextBox.x, selectedTextBox.y-1, s, true)
	case 'l':
		selectedTextBox.hide()
		selectedTextBox.placeAtXY(selectedTextBox.x+1, selectedTextBox.y, s, true)
	}
	termbox.Flush()
}

func (s *screen) prossResize(event termbox.Event) {
	switch event.Key {
	case termbox.KeyCtrlC:
		exitProgram()
	case termbox.KeyCtrlLsqBracket:
		s.changeMode(modeMove)
		placeCursorAtXY(selectedTextBox.x, selectedTextBox.y, s)
		selectedTextBox = nil
	case termbox.KeyCtrlH:
		selectedTextBox.resizeRightLeft(smaller, directionLeft, s)
	case termbox.KeyCtrlJ:
		selectedTextBox.resizeUpDown(smaller, directionDown, s)
	case termbox.KeyCtrlK:
		selectedTextBox.resizeUpDown(smaller, directionUp, s)
	case termbox.KeyCtrlL:
		selectedTextBox.resizeRightLeft(smaller, directionRight, s)
	}
	switch event.Ch {
	case 'h':
		selectedTextBox.resizeRightLeft(larger, directionLeft, s)
	case 'j':
		selectedTextBox.resizeUpDown(larger, directionDown, s)
	case 'k':
		selectedTextBox.resizeUpDown(larger, directionUp, s)
	case 'l':
		selectedTextBox.resizeRightLeft(larger, directionRight, s)
	}
	termbox.Flush()
}

func (s *screen) createAndAddTextBox() {
	tempTextBox := newTextBox()
	err := tempTextBox.placeAtXY(cursor.x, cursor.y, s, false)
	if err == nil {
		s.addTextBox(tempTextBox)
		placeCursorAtXY(cursor.x, cursor.y, s)
	}
}
