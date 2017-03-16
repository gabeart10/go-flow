package main

import (
	"errors"
	"github.com/nsf/termbox-go"
)

type textBox struct {
	width        int
	height       int
	x            int
	y            int
	shown        bool
	text         [][]rune
	border_color termbox.Attribute
	text_color   termbox.Attribute
}

type textBoxes []*textBox

func newTextBox() *textBox {
	return &textBox{
		width:        3,
		height:       3,
		x:            0,
		y:            0,
		shown:        false,
		text:         make([][]rune, 1),
		border_color: termbox.ColorBlue,
		text_color:   termbox.ColorDefault,
	}
}

func (t *textBox) placeAtXY(x, y int) error {
	w, h := termbox.Size()
	textChan := make(chan bool)
	dashChan := make(chan bool)
	pipeChan := make(chan bool)
	if t.shown == true {
		return errors.New("placeAtXY: textBox is on screen")
	} else if x+t.width-1 > w || x < 0 {
		return errors.New("placeAtXY: X is invalid")
	} else if y+t.height-1 > h || y < 0 {
		return errors.New("placeAtXY: Y is invalid")
	}
	go func(x, y int, c chan bool) {
		for i := x + 1; i < x+t.width-2; i++ {
			termbox.SetCell(i, y, '-', t.border_color, termbox.ColorDefault)
			termbox.SetCell(i, y+t.height-1, '-', t.border_color, termbox.ColorDefault)
		}
		c <- true
	}(x, y, dashChan)
	go func(x, y int, c chan bool) {
		for i := y + 1; i < y+t.height-2; i++ {
			termbox.SetCell(x, i, '|', t.border_color, termbox.ColorDefault)
			termbox.SetCell(x+t.width-1, i, '|', t.border_color, termbox.ColorDefault)
		}
		c <- true
	}(x, y, pipeChan)
	go func(x, y int, c chan bool) {
		for i := 0; i < t.height-2; i++ {
			for n := 0; n < t.width-2; n++ {
				termbox.SetCell(x+n+2, y+i+2, t.text[i][n], t.text_color, termbox.ColorDefault)
			}
		}
		c <- true
	}(x, y, textChan)
	termbox.SetCell(x, y, '+', t.border_color, termbox.ColorDefault)
	termbox.SetCell(x+t.width-1, y, '+', t.border_color, termbox.ColorDefault)
	termbox.SetCell(x, y+t.height-1, '+', t.border_color, termbox.ColorDefault)
	termbox.SetCell(x+t.width-1, y+t.height-1, '+', t.border_color, termbox.ColorDefault)
	checkDash := <-dashChan
	checkPipe := <-pipeChan
	checkText := <-textChan
	return nil
}
