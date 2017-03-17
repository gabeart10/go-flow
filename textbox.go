package main

import (
	"errors"
	"github.com/nsf/termbox-go"
)

const (
	directionUp    resizeDirection = 0
	directionDown  resizeDirection = 1
	directionRight resizeDirection = 2
	directionLeft  resizeDirection = 3
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

type cordsFunc func(int, int)

type resizeDirection int

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
	textChan := make(chan bool, 1)
	dashChan := make(chan bool, 1)
	pipeChan := make(chan bool, 1)
	if t.shown == true {
		return errors.New("placeAtXY: textBox is on screen")
	} else if x+t.width-1 > w || x < 0 {
		return errors.New("placeAtXY: X is invalid")
	} else if y+t.height-1 > h || y < 0 {
		return errors.New("placeAtXY: Y is invalid")
	}
	go func(x, y int, c chan bool) {
		for i := x + 1; i < x+t.width-1; i++ {
			termbox.SetCell(i, y, '-', t.border_color, termbox.ColorDefault)
			termbox.SetCell(i, y+t.height-1, '-', t.border_color, termbox.ColorDefault)
		}
		c <- true
	}(x, y, dashChan)
	go func(x, y int, c chan bool) {
		for i := y + 1; i < y+t.height-1; i++ {
			termbox.SetCell(x, i, '|', t.border_color, termbox.ColorDefault)
			termbox.SetCell(x+t.width-1, i, '|', t.border_color, termbox.ColorDefault)
		}
		c <- true
	}(x, y, pipeChan)
	go func(x, y int, c chan bool) {
		for i := 0; i < t.height-2; i++ {
			for n := 0; n < t.width-2; n++ {
				termbox.SetCell(x+n+1, y+i+1, t.text[i][n], t.text_color, termbox.ColorDefault)
			}
		}
		c <- true
	}(x, y, textChan)
	termbox.SetCell(x, y, '+', t.border_color, termbox.ColorDefault)
	termbox.SetCell(x+t.width-1, y, '+', t.border_color, termbox.ColorDefault)
	termbox.SetCell(x, y+t.height-1, '+', t.border_color, termbox.ColorDefault)
	termbox.SetCell(x+t.width-1, y+t.height-1, '+', t.border_color, termbox.ColorDefault)
	t.x = x
	t.y = y
	t.shown = true
	<-dashChan
	<-pipeChan
	<-textChan
	return nil
}

func (t *textBox) findAllCords(function cordsFunc) {
	for i := t.y; i < t.y+t.height; i++ {
		for n := t.x; n < t.x+t.width; n++ {
			function(n, i)
		}
	}
}

func (t *textBox) hide() {
	t.findAllCords(func(x, y int) {
		termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	})
	t.shown = false
	t.x = 0
	t.y = 0
}

func (t *textBox) resize(direction resizeDirection) {
}
