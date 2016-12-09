package main

import (
	"fmt"
	"math/rand"
	"time"

	termbox "github.com/nsf/termbox-go"
)

var (
	combo                uint     = 0
	gw                   uint     = 0
	getWildAndTough      uint     = 0
	getWildAndToughStrs  []string = []string{"Get", "Wild", "And", "Tough"}
	getWildAndToughArray []string
)

func drawLine(x, y int, str string) {
	color := termbox.ColorDefault
	backgroundColor := termbox.ColorDefault
	runes := []rune(str)

	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func drawAnimate(t time.Duration, x, y int, strs []string) {
	for _, str := range strs {
		drawLine(x, y, str)
		termbox.Flush()
		time.Sleep(t)
		y += 1
	}
}

func drawMove(x, y int, strs []string) {
	for ; x < 80; x += 1 {
		drawAnimate(10*time.Microsecond, x, y, strs)
	}
}

func drawGetWild(x, y int, strs []string) {
	drawAnimate(5*time.Millisecond, x, y, strs)
}

func draw(ev termbox.Event) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	combo += 1

	drawLine(0, 0, fmt.Sprintf("Press ESC to exit. [COMBO %d ] %v", combo, ev.Key))
	if getWildAndTough == 4 {
		drawMove(2, 1, []string{
			"      .JMMMMMMMa,                  ....        .MMM[   .MMMM    .MMM% dMM]  .MM#          .MMM`",
			"    .MMMM\"\"7\"HMMM[                .MM#         .MMM]   JMMMM   .MMMt .\"\"9`  dMMt          JMMF",
			"   JMM#'      dMM@    .JgNNNJ.  .&MMMN&,        MMM]  .MMdMN   dMM$  (gg[  .MM#    .(gNNa.MMM\\",
			"  JMM#              .MMMY\"WMMM, TUMMM\"9`        MMM] .MM^dMN  JMMD  .MMM>  JMMF  .MMMMHHMMMM#",
			" .MMM%    dMMMMMM! .MM@    dMMF  (MMF           dMM].MMF dMN .MMF   .MM#   MMM! .MMM^   .MMM]",
			" ,MMM.    \"\"\"WMMF .MMMMMMMMMMM]  MMM!           JMMbdMF  dMN.MMF    dMM]  .MM#  MMM]    .MMM`",
			" ,MMMp      .MMM% ,MM#     ...  .MM#            (MMMM#   dMNMMF    .MMM   dMM% .MMM;   .dMMF",
			"  TMMMMNggMMMMM#  .MMMN..gMM#!  dMMh.,          .MMMM`   dMMMF     JMMF  .MM#   MMMNa.&MMMM:",
			"   .TMMMMMB= TMD    TMMMMM\"'    ?HMM#           .MMM^    dMMD      MMM'  ?MMD    TMMMM\"(MM8",
		})
		getWildAndTough = 0
		termbox.Flush()
		return
	}
	switch ev.Ch {
	case 103:
		getWildAndTough = 1
		drawGetWild(2, 1, []string{
			"      .JMMMMMMMa,                  ....   ",
			"    .MMMM\"\"7\"HMMM[                .MM# ",
			"   JMM#'      dMM@    .JgNNNJ.  .&MMMN&,  ",
			"  JMM#              .MMMY\"WMMM, TUMMM\"9`",
			" .MMM%    dMMMMMM! .MM@    dMMF  (MMF     ",
			" ,MMM.    \"\"\"WMMF .MMMMMMMMMMM]  MMM!  ",
			" ,MMMp      .MMM% ,MM#     ...  .MM#      ",
			"  TMMMMNggMMMMM#  .MMMN..gMM#!  dMMh.,    ",
			"   .TMMMMMB= TMD    TMMMMM\"'    ?HMM#    ",
		})
	case 119:
		if getWildAndTough == 1 {
			drawGetWild(2, 1, []string{
				"      .JMMMMMMMa,                  ....        .MMM[   .MMMM    .MMM% dMM]  .MM#          .MMM`",
				"    .MMMM\"\"7\"HMMM[                .MM#         .MMM]   JMMMM   .MMMt .\"\"9`  dMMt          JMMF",
				"   JMM#'      dMM@    .JgNNNJ.  .&MMMN&,        MMM]  .MMdMN   dMM$  (gg[  .MM#    .(gNNa.MMM\\",
				"  JMM#              .MMMY\"WMMM, TUMMM\"9`        MMM] .MM^dMN  JMMD  .MMM>  JMMF  .MMMMHHMMMM#",
				" .MMM%    dMMMMMM! .MM@    dMMF  (MMF           dMM].MMF dMN .MMF   .MM#   MMM! .MMM^   .MMM]",
				" ,MMM.    \"\"\"WMMF .MMMMMMMMMMM]  MMM!           JMMbdMF  dMN.MMF    dMM]  .MM#  MMM]    .MMM`",
				" ,MMMp      .MMM% ,MM#     ...  .MM#            (MMMM#   dMNMMF    .MMM   dMM% .MMM;   .dMMF",
				"  TMMMMNggMMMMM#  .MMMN..gMM#!  dMMh.,          .MMMM`   dMMMF     JMMF  .MM#   MMMNa.&MMMM:",
				"   .TMMMMMB= TMD    TMMMMM\"'    ?HMM#           .MMM^    dMMD      MMM'  ?MMD    TMMMM\"(MM8",
			})
			getWildAndTough = 2
		} else {
			drawGetWild(2, 1, []string{
				"     .MMM[   .MMMM    .MMM% dMM]  .MM#          .MMM`",
				"     .MMM]   JMMMM   .MMMt .\"\"9`  dMMt          JMMF",
				"      MMM]  .MMdMN   dMM$  (gg[  .MM#    .(gNNa.MMM\\",
				"      MMM] .MM^dMN  JMMD  .MMM>  JMMF  .MMMMHHMMMM#",
				"      dMM].MMF dMN .MMF   .MM#   MMM! .MMM^   .MMM]",
				"      JMMbdMF  dMN.MMF    dMM]  .MM#  MMM]    .MMM`",
				"      (MMMM#   dMNMMF    .MMM   dMM% .MMM;   .dMMF",
				"      .MMMM`   dMMMF     JMMF  .MM#   MMMNa.&MMMM:",
				"      .MMM^    dMMD      MMM'  ?MMD    TMMMM\"(MM8",
			})
			getWildAndTough = 0
		}
	case 97:
		if getWildAndTough == 2 {
			drawGetWild(2, 1, []string{
				"      .JMMMMMMMa,                  ....        .MMM[   .MMMM    .MMM% dMM]  .MM#          .MMM`",
				"    .MMMM\"\"7\"HMMM[                .MM#         .MMM]   JMMMM   .MMMt .\"\"9`  dMMt          JMMF",
				"   JMM#'      dMM@    .JgNNNJ.  .&MMMN&,        MMM]  .MMdMN   dMM$  (gg[  .MM#    .(gNNa.MMM\\",
				"  JMM#              .MMMY\"WMMM, TUMMM\"9`        MMM] .MM^dMN  JMMD  .MMM>  JMMF  .MMMMHHMMMM#",
				" .MMM%    dMMMMMM! .MM@    dMMF  (MMF           dMM].MMF dMN .MMF   .MM#   MMM! .MMM^   .MMM]",
				" ,MMM.    \"\"\"WMMF .MMMMMMMMMMM]  MMM!           JMMbdMF  dMN.MMF    dMM]  .MM#  MMM]    .MMM`",
				" ,MMMp      .MMM% ,MM#     ...  .MM#            (MMMM#   dMNMMF    .MMM   dMM% .MMM;   .dMMF",
				"  TMMMMNggMMMMM#  .MMMN..gMM#!  dMMh.,          .MMMM`   dMMMF     JMMF  .MM#   MMMNa.&MMMM:",
				"   .TMMMMMB= TMD    TMMMMM\"'    ?HMM#           .MMM^    dMMD      MMM'  ?MMD    TMMMM\"(MM8",
				"                                             `",
				"        .J++J.                            .JJJ ",
				"       .MMMMM|                            dMM] ",
				"      .MMMMMMb     .... .....       ......MMM` ",
				"     .MMM'JMMN     -MMNMMMMMMN.  .uMMMMMMMMMF  ",
				"    .MMM! ,MMM.    MMMD`  ,MMM` .MMMY   WMMM\\ ",
				"   .MMM_...MMM]   .MM#    -MMF .MMM\\    .MM#  ",
				"  .MMMMMMMMMMMb   dMM]    MMM% .MM#     MMM]   ",
				" (MM#!~~~~~dMMN  .MMM`   .MM#  -MMMa...MMMM    ",
				"JMM#`      (MMM- (MMF    dMMF   7MMMMM#dMMF    ",
			})
			getWildAndTough = 3
		} else {
			drawGetWild(2, 1, []string{
				"                                             `",
				"        .J++J.                            .JJJ ",
				"       .MMMMM|                            dMM] ",
				"      .MMMMMMb     .... .....       ......MMM` ",
				"     .MMM'JMMN     -MMNMMMMMMN.  .uMMMMMMMMMF  ",
				"    .MMM! ,MMM.    MMMD`  ,MMM` .MMMY   WMMM\\ ",
				"   .MMM_...MMM]   .MM#    -MMF .MMM\\    .MM#  ",
				"  .MMMMMMMMMMMb   dMM]    MMM% .MM#     MMM]   ",
				" (MM#!~~~~~dMMN  .MMM`   .MM#  -MMMa...MMMM    ",
				"JMM#`      (MMM- (MMF    dMMF   7MMMMM#dMMF    ",
			})
			getWildAndTough = 0
		}
	case 116:
		if getWildAndTough == 3 {
			drawGetWild(2, 1, []string{
				"      .JMMMMMMMa,                  ....        .MMM[   .MMMM    .MMM% dMM]  .MM#          .MMM`",
				"    .MMMM\"\"7\"HMMM[                .MM#         .MMM]   JMMMM   .MMMt .\"\"9`  dMMt          JMMF",
				"   JMM#'      dMM@    .JgNNNJ.  .&MMMN&,        MMM]  .MMdMN   dMM$  (gg[  .MM#    .(gNNa.MMM\\",
				"  JMM#              .MMMY\"WMMM, TUMMM\"9`        MMM] .MM^dMN  JMMD  .MMM>  JMMF  .MMMMHHMMMM#",
				" .MMM%    dMMMMMM! .MM@    dMMF  (MMF           dMM].MMF dMN .MMF   .MM#   MMM! .MMM^   .MMM]",
				" ,MMM.    \"\"\"WMMF .MMMMMMMMMMM]  MMM!           JMMbdMF  dMN.MMF    dMM]  .MM#  MMM]    .MMM`",
				" ,MMMp      .MMM% ,MM#     ...  .MM#            (MMMM#   dMNMMF    .MMM   dMM% .MMM;   .dMMF",
				"  TMMMMNggMMMMM#  .MMMN..gMM#!  dMMh.,          .MMMM`   dMMMF     JMMF  .MM#   MMMNa.&MMMM:",
				"   .TMMMMMB= TMD    TMMMMM\"'    ?HMM#           .MMM^    dMMD      MMM'  ?MMD    TMMMM\"(MM8",
				"",
				"                                             `",
				"        .J++J.                            .JJJ        +++++++++++++,                                         .JJJ           .JJJ.  `.JJJ",
				"       .MMMMM|                            dMM]       .MMMMMMMMMMMMM`                                     `  `dMM]           gMM#   .MMMF",
				"      .MMMMMMb     .... .....       ......MMM`            dMMF       ..(J.,.    ....    ....    ....,.....  .MMM`.....     .MMM%   .MMM`",
				"     .MMM'JMMN     -MMNMMMMMMN.  .uMMMMMMMMMF      `     .MMM\\    .gMMMMMMMMp   dMMF    MMM\\  .dMMMMMMdMM]  (MMNMMMMMMN    .MM#    gMMF",
				"    .MMM! ,MMM.    MMMD`  ,MMM` .MMMY   WMMM\\            JMM#    .MMM^   /MMM. .MMM`   .MM#  .MMM=   WMMM`  MMMD`  (MMM    -MM'    MM#",
				"   .MMM_...MMM]   .MM#    -MMF .MMM\\    .MM#            .MMM]   .MMM`    .MMM! (MMF    dMM] .MMM'    JMMF  .MM#    JMMF    dMF     MM'",
				"  .MMMMMMMMMMMb   dMM]    MMM% .MM#     MMM]            .MMM`   JMM#     JMMF  MMM\\   .MMM` (MM#    .MMM:  dMM%   .MMM\\    T9     .\"D",
				" (MM#!~~~~~dMMN  .MMM`   .MM#  -MMMa...MMMM             dMMF    (MMM,...dMMD  .MMMa...MMMF  (MMMm..JMMM#  .MM#    .MM#   .ggg~   .ggm",
				"JMM#`      (MMM- (MMF    dMMF   7MMMMM#dMMF      `     .MMM>     7MMMMMMMB^    WMMMMM9MMM\\   ?MMMMM\"dMM%  JMMF    dMM]   .MM#    dMMF",
				"                                   `                                 ~`                    ....    .MM#",
				"                                                                                           JMMMNggMMM@`",
				"                                                      `                                      7\"\"\"\"\"^",
			})
			getWildAndTough = 4
		} else {
			drawGetWild(2, 1, []string{
				" +++++++++++++,                                         .JJJ           .JJJ.  `.JJJ",
				".MMMMMMMMMMMMM`                                     `  `dMM]           gMM#   .MMMF",
				"     dMMF       ..(J.,.    ....    ....    ....,.....  .MMM`.....     .MMM%   .MMM`",
				"    .MMM\\    .gMMMMMMMMp   dMMF    MMM\\  .dMMMMMMdMM]  (MMNMMMMMMN    .MM#    gMMF",
				"     JMM#    .MMM^   /MMM. .MMM`   .MM#  .MMM=   WMMM`  MMMD`  (MMM    -MM'    MM#",
				"    .MMM]   .MMM`    .MMM! (MMF    dMM] .MMM'    JMMF  .MM#    JMMF    dMF     MM'",
				"   .MMM`   JMM#     JMMF  MMM\\   .MMM` (MM#    .MMM:  dMM%   .MMM\\    T9     .\"D",
				"   dMMF    (MMM,...dMMD  .MMMa...MMMF  (MMMm..JMMM#  .MM#    .MM#   .ggg~   .ggm",
				"  .MMM>     7MMMMMMMB^    WMMMMM9MMM\\   ?MMMMM\"dMM%  JMMF    dMM]   .MM#    dMMF",
				"                ~`                    ....    .MM#",
				"                                      JMMMNggMMM@`",
				" `                                      7\"\"\"\"\"^",
			})
			getWildAndTough = 0
		}
	default:
		getWildAndTough = 0
		if ev.Key == termbox.KeyEnter {
			drawGetWild(2, 1, []string{
				"      .JMMMMMMMa,                  ....        .MMM[   .MMMM    .MMM% dMM]  .MM#          .MMM`",
				"    .MMMM\"\"7\"HMMM[                .MM#         .MMM]   JMMMM   .MMMt .\"\"9`  dMMt          JMMF",
				"   JMM#'      dMM@    .JgNNNJ.  .&MMMN&,        MMM]  .MMdMN   dMM$  (gg[  .MM#    .(gNNa.MMM\\",
				"  JMM#              .MMMY\"WMMM, TUMMM\"9`        MMM] .MM^dMN  JMMD  .MMM>  JMMF  .MMMMHHMMMM#",
				" .MMM%    dMMMMMM! .MM@    dMMF  (MMF           dMM].MMF dMN .MMF   .MM#   MMM! .MMM^   .MMM]",
				" ,MMM.    \"\"\"WMMF .MMMMMMMMMMM]  MMM!           JMMbdMF  dMN.MMF    dMM]  .MM#  MMM]    .MMM`",
				" ,MMMp      .MMM% ,MM#     ...  .MM#            (MMMM#   dMNMMF    .MMM   dMM% .MMM;   .dMMF",
				"  TMMMMNggMMMMM#  .MMMN..gMM#!  dMMh.,          .MMMM`   dMMMF     JMMF  .MM#   MMMNa.&MMMM:",
				"   .TMMMMMB= TMD    TMMMMM\"'    ?HMM#           .MMM^    dMMD      MMM'  ?MMD    TMMMM\"(MM8",
				"",
				"                                             `",
				"        .J++J.                            .JJJ        +++++++++++++,                                         .JJJ           .JJJ.  `.JJJ",
				"       .MMMMM|                            dMM]       .MMMMMMMMMMMMM`                                     `  `dMM]           gMM#   .MMMF",
				"      .MMMMMMb     .... .....       ......MMM`            dMMF       ..(J.,.    ....    ....    ....,.....  .MMM`.....     .MMM%   .MMM`",
				"     .MMM'JMMN     -MMNMMMMMMN.  .uMMMMMMMMMF      `     .MMM\\    .gMMMMMMMMp   dMMF    MMM\\  .dMMMMMMdMM]  (MMNMMMMMMN    .MM#    gMMF",
				"    .MMM! ,MMM.    MMMD`  ,MMM` .MMMY   WMMM\\            JMM#    .MMM^   /MMM. .MMM`   .MM#  .MMM=   WMMM`  MMMD`  (MMM    -MM'    MM#",
				"   .MMM_...MMM]   .MM#    -MMF .MMM\\    .MM#            .MMM]   .MMM`    .MMM! (MMF    dMM] .MMM'    JMMF  .MM#    JMMF    dMF     MM'",
				"  .MMMMMMMMMMMb   dMM]    MMM% .MM#     MMM]            .MMM`   JMM#     JMMF  MMM\\   .MMM` (MM#    .MMM:  dMM%   .MMM\\    T9     .\"D",
				" (MM#!~~~~~dMMN  .MMM`   .MM#  -MMMa...MMMM             dMMF    (MMM,...dMMD  .MMMa...MMMF  (MMMm..JMMM#  .MM#    .MM#   .ggg~   .ggm",
				"JMM#`      (MMM- (MMF    dMMF   7MMMMM#dMMF      `     .MMM>     7MMMMMMMB^    WMMMMM9MMM\\   ?MMMMM\"dMM%  JMMF    dMM]   .MM#    dMMF",
				"                                   `                                 ~`                    ....    .MM#",
				"                                                                                           JMMMNggMMM@`",
				"                                                      `                                      7\"\"\"\"\"^",
			})
			getWildAndToughArray = []string{}
		} else if ev.Key != 0 {
			drawGetWild(2, 1, []string{
				"==",
				"=========================",
				"============================================",
				"========================================================",
				"=====================================================================",
				"===========================================================",
				"======================================================",
				"========================================",
				"============================",
				"================",
			})
			drawGetWild(2, 1, []string{
				"=====================================================================",
				"=============================================================",
				"=========================================================",
				"================================================================",
				"=====================================================",
				"=============================================",
				"=============================",
				"=========================",
				"=================================",
				"====================================",
			})
			drawGetWild(2, 1, []string{
				"=========================================",
				"=================================",
				"===================================",
				"============================================",
				"======================================================",
				"===========================================",
				"==================================================",
				"==================================================================",
				"========================================================================",
				"=====================================================================",
			})
		} else {
			getWildAndToughArray = append(getWildAndToughArray, getWildAndToughStrs[gw%4])
			gw += 1
			rand.Seed(time.Now().UnixNano())
			i := rand.Intn(4) + 2
			drawLine(i, i, fmt.Sprintf("%v", getWildAndToughArray))
		}
	}

	termbox.Flush()
}

func pollEvent() {
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
