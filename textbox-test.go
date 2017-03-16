package main

import (
	"github.com/nsf/termbox-go"
	"os"
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
	err := textb.placeAtXY(7, 10)
	if err != nil {
		print(err)
	}
	for {
	}
}
