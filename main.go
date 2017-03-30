package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

func main() {
	termbox.Init()
	screena := newScreen()
	termbox.Flush()
	textb := newTextBox()
	screena.addTextBox(textb)
	textb.text[0] = append(textb.text[0], ' ')
	textb.placeAtXY(15, 11, screena)
	textb.text[0][0] = 'a'
	termbox.Flush()
	time.Sleep(1 * time.Second)
	for i := 0; i < 3; i++ {
		textb.resizeUpDown(larger, directionDown, screena)
		termbox.Flush()
	}
	time.Sleep(1 * time.Second)
	for i := 0; i < 3; i++ {
		textb.resizeRightLeft(larger, directionRight, screena)
		termbox.Flush()
	}
	time.Sleep(1 * time.Second)
	for i := 0; i < 3; i++ {
		textb.resizeRightLeft(smaller, directionRight, screena)
		termbox.Flush()
	}
	placeCursorAtXY(15, 11, screena)
	termbox.Flush()
	screena.startEventWatcher()
}
