package main

import "C"

import (
	"github.com/go-vgo/robotgo"
)

// Util
func GotoCStr(s string) *C.char {
	return C.CString(s)
}

func CtoGoStr(ch *C.char) string {
	return C.GoString(ch)
}

//export Sleep
func Sleep(s int) {
	robotgo.Sleep(s)
}

//export MilliSleep
func MilliSleep(ms int) {
	robotgo.MilliSleep(ms)
}

// Mouse API

//export SetMouseSleep
func SetMouseSleep(ms int) {
	robotgo.MouseSleep = ms // Default is 10ms
}

//export GetMousePositionX
func GetMousePositionX() int {
	x, _ := robotgo.Location()
	return x
}

//export GetMousePositionY
func GetMousePositionY() int {
	_, y := robotgo.Location()
	return y
}

//export DragMouseSmooth
func DragMouseSmooth(x, y int, low, high float64, mouseDelay int) {
	robotgo.DragSmooth(x, y, low, high, mouseDelay)
}

//export MoveMouse
func MoveMouse(x, y int) {
	robotgo.Move(x, y)
}

//export MoveMouseSmooth
func MoveMouseSmooth(x, y int, low, high float64, mouseDelay int) {
	robotgo.MoveSmooth(x, y, low, high, mouseDelay)
}

//export MoveMouseRelative
func MoveMouseRelative(x, y int) {
	robotgo.MoveRelative(x, y)
}

//export MoveMouseRelativeSmooth
func MoveMouseRelativeSmooth(x, y int, low, high float64, mouseDelay int) {
	robotgo.MoveSmoothRelative(x, y, low, high, mouseDelay)
}

//export ScrollMouse
func ScrollMouse(x, y int) {
	robotgo.Scroll(x, y)
}

//export MouseClick
func MouseClick(btn *C.char) {
	robotgo.Click(CtoGoStr(btn))
}

//export MouseToggle
func MouseToggle(btn, state *C.char) {
	robotgo.Toggle(CtoGoStr(btn), CtoGoStr(state))
}

// Keyboard API

//export TypeString
func TypeString(str *C.char) {
	robotgo.TypeStr(CtoGoStr(str))
}

//export KeyPress
func KeyPress(key *C.char) {
	robotgo.KeyPress(CtoGoStr(key))
}

//export KeyDown
func KeyDown(key *C.char) {
	robotgo.KeyDown(CtoGoStr(key))
}

//export KeyUp
func KeyUp(key *C.char) {
	robotgo.KeyUp(CtoGoStr(key))
}

// Main
func main () {}
