package main

import (
	"fmt"
	"time"

	termbox "github.com/nsf/termbox-go"
)

var (
	combo           uint = 0
	getWildAndTough []string
)

func drawLine(x, y int, str string) {
	color := termbox.ColorDefault
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func drawAnimate(x, y int, strs []string) {
	for _, str := range strs {
		drawLine(x, y, str)
		termbox.Flush()
		time.Sleep(20 * time.Millisecond)
	}
}

func drawMove(x, y int, str string) {
	n := y + 25
	for ; y < n; n -= 1 {
		drawLine(x, n, str)
		termbox.Flush()
		time.Sleep(10 * time.Millisecond)
	}
}

func drawGetWild(x, y int, strs []string) {
	for _, str := range strs {
		drawLine(x, y, str)
		termbox.Flush()
		time.Sleep(20 * time.Millisecond)
		y += 1
	}
}

func draw(ev termbox.Event) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	combo += 1

	drawLine(0, 0, fmt.Sprintf("Press ESC to exit. [COMBO %d ]", combo))
	switch ev.Ch {
	case 103:
		getWildAndTough = append(getWildAndTough, "GET")
		drawLine(2, 1, fmt.Sprintf("%v ======== %d COMMBO", getWildAndTough, combo))
	case 119:
		getWildAndTough = append(getWildAndTough, "WILD")
		drawLine(3, 2, fmt.Sprintf("%v ======== %d COMMBO", getWildAndTough, combo))
	case 97:
		getWildAndTough = append(getWildAndTough, "AND")
		drawLine(4, 3, fmt.Sprintf("%v ======== %d COMMBO", getWildAndTough, combo))
	case 116:
		getWildAndTough = append(getWildAndTough, "TOUGH!!")
		drawLine(5, 4, fmt.Sprintf("%v ======== %d COMMBO", getWildAndTough, combo))
	default:
		getWildAndTough = []string{}
		if ev.Key == termbox.KeyEnter {
			//drawMove(2, 1, "==============================================================================================")
			drawGetWild(2, 1, []string{
				"　 　,,,,,lllllllllllllllll,,,,, 　 　　 lllllllllllllllllllllllllllllllll,　,llllllllllllllllllllllllllllllllllll 　 　　 　'lllllllll, 　 　 ,llllllllll 　　　 lllllllll　 lllllllll, 　　lllllllll",
				"　,,lllllllllllll'''''''lllllllllllll,, 　 　llllllllllllllllllllllllllllllllll 　llllllllllllllllllllllllllllllllllll　 　 　 　 llllllllll　 　 ,llllllllllll　　　,lllllllll 　 lllllllll　 　llllllllll",
				"　llllllll''　　　 　''llllllllll　 　llllllllll　 　 　 　 　 　　 　　llllllllll　　　　 　　 　　　llllllll,　　 llllll''llllll,　　 llllllll' 　　lllllllll　 　lllllllll",
				",lllllllll' 　　　 　 　 　　 　　lllllllll,,,,,,,,,,,,,,,,,,,,,, 　　　 　 lllllllllll　　 　　　　　　　 'lllllllll 　 llllll' 'lllllll 　 llllllll'　 　llllllllll　　 llllllllll",
				"llllllllll　 　 　llllllllllllllllll 　　llllllllllllllllllllllllllllllll　 　 　　 lllllllllll　 　 　　　　　 　　'llllllll,　,llllll'　'llllll,　,llllllll' 　　 lllllllll　　 llllllllll",
				"llllllllll 　 　　llllllllllllllllll 　　llllllllll 　 　　 　　 　 　　　 lllllllll 　 　 　　 　　　 　　llllllll ,llllll' 　 llllll, llllllll' 　 　 llllllllll　　 lllllllll",
				"　llllllll,, 　 　 　 ,,lllllllll　 　llllllllll 　 　 　　 　 　　　　 lllllllll　 　 　　　 　 　 　　 lllllll,llllll' 　　llllllllllllll'　 　 　lllllllll　 　llllllllll",
				"　'llllllllllll,,,,,,,,,,,lllllllllllllll　　 llllllllllllllllllllllllllllllllll, 　　　　lllllllllll 　　　 　　 　 　　 　 lllllllllllll 　　 'llllllllllll 　 　 　lllllllll　　 lllllllllllllll",
				"　　'''llllllllllllllllllll''' 'lllllll 　　lllllllllllllllllllllllllllllllllll 　 　　 lllllllllll 　 　 　 　 　 　　　 　llllllllll　　 　'lllllllll' 　 　 　 lllllllll　　 lllllllllllllll",
			})
		} else {
			drawAnimate(2, 1, []string{
				"==",
				"===",
				"=====",
				"========",
				"==========",
				"===================",
				"==============",
				"==========",
				"========",
				"======",
			})
		}
	}

	termbox.Flush()
}

func pollEvent() {
	tm := termbox.Event{}
	draw(tm)
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc,
				termbox.KeyCtrlC,
				termbox.KeyCtrlD:
				return
			default:
				draw(ev)
			}
		default:
			draw(ev)
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	pollEvent()
}
