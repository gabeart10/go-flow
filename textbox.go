package main

import (
	"errors"
	"github.com/nsf/termbox-go"
)

const (
	directionUp resizeOption = iota
	directionDown
	directionRight
	directionLeft
	larger
	smaller
)

type textBox struct {
	width        int
	height       int
	x            int
	y            int
	shown        bool
	text         [][]rune
	allCords     []cords
	border_color termbox.Attribute
	text_color   termbox.Attribute
}

var borderBox = &textBox{
	width:  0,
	height: 0,
}

type textBoxes []*textBox

type cords [2]int

type cordsFunc func(int, int)

type cordsFuncOp func(int, int, int)

type resizeOption int

func newTextBox() *textBox {
	returnBox := &textBox{
		width:        3,
		height:       3,
		x:            0,
		y:            0,
		shown:        false,
		text:         make([][]rune, 1),
		allCords:     make([]cords, 9),
		border_color: termbox.ColorBlue,
		text_color:   termbox.ColorDefault,
	}
	returnBox.text[0] = append(returnBox.text[0], 0x0)
	return returnBox
}

func (t *textBox) placeAtXY(x, y int) error {
	w, h := termbox.Size()
	h--
	textChan := make(chan bool, 1)
	dashChan := make(chan bool, 1)
	pipeChan := make(chan bool, 1)
	cordsChan := make(chan bool, 1)
	if t.shown == true {
		return errors.New("placeAtXY: textBox is on screen")
	} else if x+t.width-1 > w || x < 0 {
		return errors.New("placeAtXY: X is invalid")
	} else if y+t.height-1 > h || y < 0 {
		return errors.New("placeAtXY: Y is invalid")
	}
	t.x = x
	t.y = y
	t.shown = true
	go t.updateCords(cordsChan)

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
				if t.text[i][n] == 0x0 {
					termbox.SetCell(x+n+1, y+i+1, ' ', t.text_color, termbox.ColorDefault)
				} else {
					termbox.SetCell(x+n+1, y+i+1, t.text[i][n], t.text_color, termbox.ColorDefault)
				}
			}
		}
		c <- true
	}(x, y, textChan)
	termbox.SetCell(x, y, '+', t.border_color, termbox.ColorDefault)
	termbox.SetCell(x+t.width-1, y, '+', t.border_color, termbox.ColorDefault)
	termbox.SetCell(x, y+t.height-1, '+', t.border_color, termbox.ColorDefault)
	termbox.SetCell(x+t.width-1, y+t.height-1, '+', t.border_color, termbox.ColorDefault)
	<-dashChan
	<-pipeChan
	<-textChan
	<-cordsChan
	return nil
}

func (t *textBox) updateCords(done chan bool) {
	t.allCords = make([]cords, t.width*t.height)
	t.findAllCordsOp(func(x, y, c int) {
		t.allCords[c] = [2]int{x, y}
	})
	done <- true
}

func (t *textBox) findAllCords(function cordsFunc) {
	for i := t.y; i < t.y+t.height; i++ {
		for n := t.x; n < t.x+t.width; n++ {
			function(n, i)
		}
	}
}
func (t *textBox) findAllCordsOp(function cordsFuncOp) {
	counter := 0
	for i := t.y; i < t.y+t.height; i++ {
		for n := t.x; n < t.x+t.width; n++ {
			function(n, i, counter)
			counter++
		}
	}
}

func (t *textBox) hide() {
	t.findAllCords(func(x, y int) {
		termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	})
	t.shown = false
}

func (t *textBox) subColliding(currentBox *textBox, found chan *textBox) {
	for _, currentCord := range t.allCords {
		for _, compCord := range currentBox.allCords {
			if currentCord == compCord {
				found <- currentBox
			}
		}
	}
	found <- nil
}

func (s *screen) checkIfColliding(t *textBox) *textBox {
	found := make(chan *textBox, 1)
	w, h := termbox.Size()
	h--
	sent := 0
	if t.x+t.width-1 > w || t.x < 0 {
		return borderBox
	} else if t.y+t.height-1 > h || t.y < 0 {
		return borderBox
	}
	for _, currentBox := range s.boxes {
		if t != currentBox && currentBox != nil {
			go t.subColliding(currentBox, found)
			sent++
		}
	}
	for {
		if sent == 0 {
			return nil
		}
		isFound := <-found
		if isFound != nil {
			return isFound
		}
		sent--
	}
}

func (t *textBox) resizeUpDown(largerSmaller resizeOption, upDown resizeOption, s *screen) error {
	if t.shown == false {
		return errors.New("resize: textBox is not shown")
	}
	t.hide()
	if largerSmaller == larger {
		if upDown == directionUp {
			t.y--
		}
		t.height++
		if s.checkIfColliding(t) != nil {
			if upDown == directionUp {
				t.y++
			}
			t.height--
			return errors.New("resize: Object in way")
		}
		newText := make([][]rune, t.height-2)
		for i := 0; i < t.height-3; i++ {
			for n := 0; n < t.width-2; n++ {
				newText[i] = append(newText[i], t.text[i][n])
			}
		}
		for i := 0; i < t.width-2; i++ {
			newText[t.height-3] = append(newText[t.height-3], 0x0)
		}
		t.text = newText
	} else if largerSmaller == smaller {
		if t.height <= 3 {
			return errors.New("resize: Textbox too small")
		}
		for _, symbol := range t.text[t.height-3] {
			if symbol != 0x0 {
				return errors.New("resize: Text in way")
			}
		}
		t.height--
		if upDown == directionUp {
			t.y++
		}
		newText := make([][]rune, t.height-2)
		for i := 0; i < t.height-2; i++ {
			for n := 0; n < t.width-2; n++ {
				newText[i] = append(newText[i], t.text[i][n])
			}
		}
		t.text = newText
	}
	t.placeAtXY(t.x, t.y)
	return nil
}

func (t *textBox) resizeRightLeft(largerSmaller resizeOption, leftRight resizeOption, s *screen) error {
	if t.shown == false {
		return errors.New("resize: Textbox is not shown")
	}
	t.hide()
	if largerSmaller == larger {
		t.width++
		if leftRight == directionLeft {
			t.x--
		}
		if s.checkIfColliding(t) != nil {
			t.width--
			if leftRight == directionLeft {
				t.x++
			}
			errors.New("resize: Object in way")
		}
		newText := make([][]rune, t.height-2)
		x := 0
		y := 0
		for i := 0; i < t.height-2; i++ {
			for n := 0; n < t.width-3; n++ {
				if x < t.width-2 {
					newText[y] = append(newText[y], t.text[i][n])
					x++
				} else {
					x = 0
					y++
				}
			}
		}
		for i := x; i < t.width-2; i++ {
			newText[y] = append(newText[y], 0x0)
		}
		for i := y; i < t.height-2; i++ {
			for n := 0; n < t.width-2; n++ {
				newText[i] = append(newText[i], 0x0)
			}
		}
		t.text = newText
	}
	t.placeAtXY(t.x, t.y)
	return nil
}
