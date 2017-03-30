package main

import (
	"github.com/nsf/termbox-go"
	"os"
)

var selectedTextBox = nil

func (s *screen) startEventWatcher() {
	for {
		event := termbox.PollEvent()
		switch s.screenMode {
		case modeMove:
			s.prossMove(event)
		case modeResize:
			//s.prossResize(event)
		case modeMoveBox:
			//s.prossMoveBox(event)
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
		termbox.Flush()
	case 'k':
		placeCursorAtXY(cursor.x, cursor.y-1, s)
		termbox.Flush()
	case 'j':
		placeCursorAtXY(cursor.x, cursor.y+1, s)
		termbox.Flush()
	case 'l':
		placeCursorAtXY(cursor.x+1, cursor.y, s)
		termbox.Flush()
	case 'm':
		collData := checkCursorColl(s)
		if collData != nil {
			s.changeMode(modeMoveBox)
			selectedTextBox = collData
		}
	}
}
