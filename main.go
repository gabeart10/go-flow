package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	termbox.Init()
	screena := newScreen()
	placeCursorAtXY(15, 11, screena)
	termbox.Flush()
	screena.startEventWatcher()
}
