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
	texta := newTextBox()
	screena.addTextBox(textb)
	screena.addTextBox(texta)
	textb.text[0] = append(textb.text[0], ' ')
	texta.text[0] = append(textb.text[0], ' ')
	textb.placeAtXY(15, 6)
	texta.placeAtXY(15, 10)
	result := screena.checkIfColliding(textb)
	termbox.Flush()
	time.Sleep(1 * time.Second)
	if result == true {
		termbox.Close()
		os.Exit(3)
	}

	for {
	}
}
