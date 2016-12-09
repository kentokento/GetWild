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
				"                  ...........",
				"             ..gMMMMMMMMMMMMMMNa.                                                                 .MMMMMMM]         .MMMMMMMF         .MMMMMMMt   MMMMMM#      .MMMMMM\\                      JMMMMMM`",
				"           .MMMMMMMMMMMMMMMMMMMMMN,                                     MMMMMM#                    MMMMMMMF        .MMMMMMMMF         dMMMMMMt   .MMMMMMF      MMMMMM#                      .MMMMMMF",
				"        ..MMMMMMMMMM\"\"\"\"\"MMMMMMMMMM,                                   .MMMMMM%                    MMMMMMMF       .MMMMMMMMMF        JMMMMMM$    dMMMMMM`     .MMMMMM]                      -MMMMMM\\",
				"       .MMMMMMMM\"`         ?MMMMMMMN.                                  dMMMMM#                     dMMMMMMF       dMMMMMMMMMF       -MMMMMM$                  gMMMMMM                       MMMMMM#",
				"     .dMMMMMMM^             .MMMMMMM)          ..JgNMMNNgJ.        ....MMMMMMb....                 JMMMMMMF      JMMMMFMMMMMF      .MMMMMM$     ........     .MMMMMMF          ..JgNNNaJ.  .MMMMMM]",
				"    .MMMMMMMF                7\"\"\"\"\"\"!       .+MMMMMMMMMMMMMMm.    .MMMMMMMMMMMMMM>                 ,MMMMMM@     .MMMM# MMMMMF     .MMMMMM$      MMMMMM#      -MMMMMM!       .JMMMMMMMMMMMN,dMMMMMM",
				"    dMMMMMMF                              .MMMMMMMMMMMMMMMMMMN,   (MMMMMMMMMMMMM#                  .MMMMMM#    .MMMMM! MMMMMF    .MMMMMMD      .MMMMMMt      MMMMMMF      .dMMMMMMMMMMMMMMMMMMMMMF",
				"   -MMMMMM#                              JMMMMMM@^    .UMMMMMMN      .MMMMMM>                      .MMMMMM#   .MMMMM^  MMMMMF   .MMMMMMF       dMMMMMM      .MMMMMM\\     .MMMMMMM#\"!  ?TMMMMMMMMM!",
				"  .MMMMMMM\\         .MMMMMMMMMMMMMM`   .MMMMMM#!        HMMMMMM.     dMMMMM#                        MMMMMM#   MMMMMD   MMMMMF  .MMMMMMF       .MMMMMMF      dMMMMM#     .MMMMMMB`       .MMMMMMMF",
				"  -MMMMMM#          dMMMMMMMMMMMMMF    dMMMMMM(.........dMMMMMM`    .MMMMMM%                        MMMMMM#  JMMMMF    MMMMM]  MMMMMMF        JMMMMMM!     .MMMMMMF    .MMMMMMF          dMMMMMM%",
				"  dMMMMMMF         .MMMMMMMMMMMMMM\\   .MMMMMMMMMMMMMMMMMMMMMMM#     JMMMMM#                         dMMMMM# .MMMM@     MMMMM] JMMMMMF        .MMMMMMF      JMMMMMM`   .MMMMMM#           dMMMMM#",
				"  MMMMMMMb                 dMMMMM#    dMMMMMMMMMMMMMMMMMMMMMMM]    .MMMMMMF                         JMMMMM#.MMMM#`     MMMMM]-MMMMMF         -MMMMMM\\     .MMMMMMF    -MMMMMM]          .MMMMMMF",
				"  dMMMMMMN.              .MMMMMMM]   .MMMMMMF                      -MMMMMM`                         -MMMMMNMMMMM'     .MMMMMNMMMMMF          MMMMMM#      -MMMMMM:    dMMMMMM)          dMMMMMM`",
				"  -MMMMMMMN,           .JMMMMMMMM    .MMMMMMb         ........     MMMMMMF                          .MMMMMMMMMM\\      .MMMMMMMMMMF          .MMMMMM]      MMMMMM#     MMMMMMMb        .dMMMMMMF",
				"   UMMMMMMMMNa.......+MMMMMMMMMMF     MMMMMMMm.     .MMMMMMM$     .MMMMMM]                          .MMMMMMMMMF       .MMMMMMMMMF           dMMMMMM`     .MMMMMM%     JMMMMMMMa.   ..uMMMMMMMM>",
				"    ?MMMMMMMMMMMMMMMMMMMM#4MMMMM!     ,MMMMMMMMMMMMMMMMMMM@`      dMMMMMMMMMM                        MMMMMMMMF        .MMMMMMMMF           .MMMMMMF      dMMMMM#       WMMMMMMMMMMMMMMMMMMMMM#",
				"      TMMMMMMMMMMMMMMMM@! ,MMMMF        TMMMMMMMMMMMMMMM\"`        ?MMMMMMMMMF                        MMMMMMM@         .MMMMMMMF            (MMMMMM!     .MMMMMMF        TMMMMMMMMMMMMD.MMMMMM%",
				"        .\"\"MMMMMMMH9\"`    ,\"\"\"\"'          .\"\"MMMMMM\"\"=              ?\"\"\"\"\"\"\"!                        ?\"\"\"\"\"\"`         .\"\"\"\"\"\"\"             \"\"\"\"\"\"\"      ,\"\"\"\"\"\"`          ?\"MMMMM\"\"!  ,\"\"\"\"\"\"",
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
