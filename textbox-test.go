package main

import (
	"github.com/nsf/termbox-go"
	"os"
	"time"
)

func main() {
	termbox.Init()
	go func() {
		for {
			var event = termbox.PollEvent()
			if event.Key == termbox.KeyCtrlC {
				termbox.Close()
				os.Exit(3)
			}
		}
	}()
	screena := newScreen()
	textb := newTextBox()
	screena.addTextBox(textb)
	textb.text[0] = append(textb.text[0], ' ')
	textb.placeAtXY(15, 11)
	textb.text[0][0] = 'a'
	termbox.Flush()
	time.Sleep(1 * time.Second)
	for i := 0; i < 3; i++ {
		textb.resizeUpDown(larger, directionDown, screena)
		termbox.Flush()
	}
	time.Sleep(1 * time.Second)
	for i := 0; i < 3; i++ {
		textb.resizeRightLeft(larger, directionLeft, screena)
		termbox.Flush()
	}
	for {
	}
}
