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
	termbox.Flush()
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		textb.resizeUpDown(larger, directionDown, screena)
		termbox.Flush()
	}
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		textb.resizeUpDown(smaller, directionDown, screena)
		termbox.Flush()
	}
	for {
	}
}
