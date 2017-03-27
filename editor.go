package main

import (
	"github.com/nsf/termbox-go"
)

var cursor = &textBox{
	width:    1,
	height:   1,
	x:        0,
	y:        0,
	shown:    true,
	allCords: make([]cords, 1),
}
