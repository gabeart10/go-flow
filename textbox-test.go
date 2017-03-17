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
	textb := newTextBox()
	textb.text[0] = append(textb.text[0], ' ')
	textb.placeAtXY(20, 20)
	termbox.Flush()
	time.Sleep(1 * time.Second)
	textb.hide()
	termbox.Flush()
	for {
	}
}
