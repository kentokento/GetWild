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
		time.Sleep(5 * time.Millisecond)
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
		drawGetWild(2, 1, []string{
			"                      ...(ggNNNNNggJ...",
			"                  .JMMMMMMMMMMMMMMMMMMMMa,",
			"               .MMMMMMMMMMMMMMMMMMMMMMMMMMNx                                                   gNNNNNNNK",
			"            .JMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN,                                                .MMMMMMMMF",
			"          .dMMMMMMMMMMMMY\"\"7??7\"\"MMMMMMMMMMMMM,                                               dMMMMMMMM`",
			"         .MMMMMMMMMM#\"             ?MMMMMMMMMMN                  `  `                        .MMMMMMMMF",
			"       .MMMMMMMMMM@`                 MMMMMMMMMM;  `  `  `  `  `  ..........                  (MMMMMMMM\\",
			"      .MMMMMMMMM#'                   ,MMMMMMMMM%            ..gMMMMMMMMMMMMMNa,  `   `  JMMMMMMMMMMMMMMMMMM%",
			"     .MMMMMMMMMD                      `````````           .MMMMMMMMMMMMMMMMMMMMN,      .MMMMMMMMMMMMMMMMMM#",
			"    .MMMMMMMMMF                                        .(MMMMMMMMMMMMMMMMMMMMMMMMp     -MMMMMMMMMMMMMMMMMMF",
			"   .MMMMMMMMM#                                       `.MMMMMMMMMMB\"7?7\"WMMMMMMMMMM,    7\"\"\"TMMMMMMMM@\"\"\"\"\"`",
			"   -MMMMMMMMM\\             ...................    `  JMMMMMMMM#'         WMMMMMMMMN        JMMMMMMMM'",
			"  .MMMMMMMMMF             .MMMMMMMMMMMMMMMMMMF      JMMMMMMMM$           .MMMMMMMMM       .MMMMMMMM#            `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `",
			"  .MMMMMMMMM%             MMMMMMMMMMMMMMMMMMM\\     -MMMMMMMM$             MMMMMMMMM       .MMMMMMMM%",
			"  dMMMMMMMMM`         `  .MMMMMMMMMMMMMMMMMM#     .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM#       dMMMMMMM#",
			"  MMMMMMMMM#             dMMMMMMMMMMMMMMMMMM]     dMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMF    ` .MMMMMMMMF",
			"  MMMMMMMMMN                       dMMMMMMMM`    .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM`      JMMMMMMMM!",
			"  MMMMMMMMMM;    `  `            .dMMMMMMMMF     .MMMMMMMM]                             .MMMMMMMMF            `   `    `    `    `    `    `    `    `    `    `    `    `    `    `    `    `    `",
			"  dMMMMMMMMMN,                  .MMMMMMMMMM>     JMMMMMMMM]                          `  -MMMMMMMM\\         `",
			"  ,MMMMMMMMMMMe.             ..MMMMMMMMMMM#      JMMMMMMMMb           .Jgggggggg:       MMMMMMMM#",
			"   ?MMMMMMMMMMMMNJ.. `  ...JMMMMMMMMMMMMMM%      ,MMMMMMMMMm,       ..MMMMMMMMM$       .MMMMMMMMF      `            `    `    `    `    `    `    `    `    `    `    `    `    `    `    `    `    `",
			"    ?MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM        UMMMMMMMMMMMNgggNMMMMMMMMMM#'        dMMMMMMMMMNNNN]",
			"     .WMMMMMMMMMMMMMMMMMMMMMMMMM\" -MMMMMMF         TMMMMMMMMMMMMMMMMMMMMMMM#^          MMMMMMMMMMMMMM!      `  `",
			"       .TMMMMMMMMMMMMMMMMMMMMB^   (MMMMMM!      `    TMMMMMMMMMMMMMMMMMMM\"`          ` ,MMMMMMMMMMMM#",
			"          .\"WMMMMMMMMMMMMB\"`      ?\"\"\"\"\"\"              .TWMMMMMMMMMMY\"^                  _\"\"HMMMMMMH\\",
		})
	case 119:
		getWildAndTough = append(getWildAndTough, "WILD")
		drawGetWild(2, 1, []string{
			" (gggggggg;           .gggggggge            .gggggggg`   .NNNNNNN]       (ggggggg~                          .gggggggc",
			" JMMMMMMMM]          .MMMMMMMMM#           .MMMMMMMM^    MMMMMMMM`      .MMMMMMM#                           MMMMMMMM!",
			" -MMMMMMMM]          MMMMMMMMMM#          .MMMMMMMM%    .MMMMMMMF       .MMMMMMM]                          .MMMMMMMF",
			" ,MMMMMMMM]         JMMMMMMMMMM#         .MMMMMMMMt     dMMMMMMM\\       MMMMMMMM`                          JMMMMMMM%",
			" .MMMMMMMM]        .MMMMMMMMMMM#         MMMMMMMM$                     .MMMMMMMF                          .MMMMMMM#",
			"  MMMMMMMMF       .MMMMM#dMMMMM#        dMMMMMMMD       .......        dMMMMMMM:              ....(...    -MMMMMMMF",
			"  MMMMMMMMF      .MMMMMM^dMMMMM#       (MMMMMMMF      .MMMMMMMM       .MMMMMMM#           .gMMMMMMMMMMMa. MMMMMMMM`",
			"  dMMMMMMMF      MMMMMMF MMMMMM#    ` .MMMMMMMF       .MMMMMMMF       -MMMMMMM%        .JMMMMMMMMMMMMMMMN(MMMMMMMF",
			"  JMMMMMMMF     JMMMMMF  MMMMMM@     .MMMMMMMF        dMMMMMMM!       MMMMMMM#        -MMMMMMMMMMMMMMMMMMMMMMMMMM\\",
			"  (MMMMMMM@    .MMMMM#   MMMMMM@    .MMMMMMMF        .MMMMMMM#    `  .MMMMMMMF      .MMMMMMMMM@=`   _TMMMMMMMMMM#",
			"  ,MMMMMMM@   .MMMMMM!   MMMMMM@   .MMMMMMMF         JMMMMMMM%       dMMMMMMM!     .MMMMMMMMF         .MMMMMMMMM]",
			"  .MMMMMMM#  .MMMMMM^    MMMMMMF  .MMMMMMMF         .MMMMMMM#       .MMMMMMMF     .MMMMMMMMt           -MMMMMMMM`",
			"   MMMMMMM#  dMMMMMF     MMMMMMF  dMMMMMMF          -MMMMMMMF       JMMMMMMM\\     dMMMMMMMF            ,MMMMMMMF",
			"  `MMMMMMM@ (MMMMMF      MMMMMMF -MMMMMM@        `  MMMMMMMM!    ` .MMMMMMM#     .MMMMMMM#             JMMMMMMM>        `   `   `   `   `   `   `   `   `   `   `   `   `   `   `   `   `   `   `   `",
			"   dMMMMMMF.MMMMM@       MMMMMMF.MMMMMM@           .MMMMMMMF       .MMMMMMMF     dMMMMMMMF            .MMMMMMM#",
			"   dMMMMMMNMMMMMM`       MMMMMMNMMMMMM@            dMMMMMMM\\       dMMMMMMM`     MMMMMMMM]            dMMMMMMMt",
			"   (MMMMMMMMMMMM^        MMMMMMMMMMMM#         `  .MMMMMMM#       .MMMMMMMF      MMMMMMMMb          .JMMMMMMMM        `",
			"   ,MMMMMMMMMMM$         MMMMMMMMMMM#`            -MMMMMMMF       JMMMMMMM\\      MMMMMMMMM,        .MMMMMMMMMF            `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `  `",
			"   .MMMMMMMMMMF      `   MMMMMMMMMMM`      `      MMMMMMMM`    ` .MMMMMMM#    `  -MMMMMMMMMNgJ.(gMMMMMMMMMMMM!     `",
			"    MMMMMMMMM@           MMMMMMMMMM`             .MMMMMMMF       .MMMMMMM]        ?MMMMMMMMMMMMMMMMMBMMMMMMM#",
			"    MMMMMMMM#            MMMMMMMMM!           `  dMMMMMMM\\       MMMMMMMM`         ,MMMMMMMMMMMMMM@!.MMMMMMM%          `",
			"    7\"\"\"\"\"\"\"!            \"\"\"\"\"\"\"\"!               \"\"\"\"\"\"\"\"    `  .\"\"\"\"\"\"\"=            .THMMMMMMW\"^   ?\"\"\"\"\"\"\"",
		})
	case 97:
		getWildAndTough = append(getWildAndTough, "AND")
		drawGetWild(2, 1, []string{
			"                             .............                                                                                     ..........",
			"                            .MMMMMMMMMMMM]                                                                                     dMMMMMMMMF",
			"                           .MMMMMMMMMMMMMb                                                                                    .MMMMMMMMM:",
			"                          .MMMMMMMMMMMMMMN                                                                                    JMMMMMMMM#",
			"                         .MMMMMMMMMMMMMMMM-                                                                                  .MMMMMMMMM%",
			"                        (MMMMMMMMMMMMMMMMM]                                                                                  .MMMMMMMMM",
			"                       JMMMMMMMM@MMMMMMMMMN           `     `    `  `   `  ` ........     `  `                 ........      MMMMMMMMMF",
			"                 `   .dMMMMMMMMF MMMMMMMMMN.                .MMMMMMMM#   .&MMMMMMMMMMMNa,                 ..gMMMMMMMMMMMN,  .MMMMMMMMM!",
			"                    .MMMMMMMMMF  JMMMMMMMMM|    `  `     `  JMMMMMMMM%.JMMMMMMMMMMMMMMMMMN.      `  `   .MMMMMMMMMMMMMMMMMN,dMMMMMMMMF",
			"                   .MMMMMMMMMF   .MMMMMMMMMb               .MMMMMMMMMJMMMMMMMMMMMMMMMMMMMMN           .MMMMMMMMMMMMMMMMMMMMNMMMMMMMMM%",
			"                  .MMMMMMMMMF    .MMMMMMMMMN               .MMMMMMMMMMMMMMH\"\"WMMMMMMMMMMMMM-        .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM#",
			"     `   `  `  ` .MMMMMMMMMF      MMMMMMMMMM.         `    dMMMMMMMMMMB^        TMMMMMMMMMM:       .MMMMMMMMMM@^       TMMMMMMMMMMMMF        `   `   `   `   `   `   `   `   `   `   `   `   `   `",
			"                .MMMMMMMMMD       dMMMMMMMMM]             .MMMMMMMMMM^           MMMMMMMMMM       .MMMMMMMMM@`          .MMMMMMMMMMM`",
			"               .MMMMMMMMM$        JMMMMMMMMMb             JMMMMMMMMM`            MMMMMMMMMF      .MMMMMMMMMD             (MMMMMMMMMF",
			"              (MMMMMMMMMt         ,MMMMMMMMMN        `   .MMMMMMMMM%            .MMMMMMMMM!     .MMMMMMMMMF              .MMMMMMMMM\\",
			"             JMMMMMMMMM\\          .MMMMMMMMMM;           -MMMMMMMM#             dMMMMMMMM#      dMMMMMMMM#               -MMMMMMMM#",
			"           .dMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM]           MMMMMMMMM%            .MMMMMMMMM%     .MMMMMMMMM\\               dMMMMMMMM]",
			"          .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM#       `  .MMMMMMMM#             JMMMMMMMM#      dMMMMMMMMM               .MMMMMMMMM",
			"         .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM.         dMMMMMMMMF         `  .MMMMMMMMMF      MMMMMMMMMN              .MMMMMMMMMF",
			"        .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM[        .MMMMMMMMM!            .MMMMMMMMM!      MMMMMMMMMM.            .MMMMMMMMMM:",
			"       .MMMMMMMMMM^?????????????????MMMMMMMMMMb        (MMMMMMMMF             MMMMMMMMMF       MMMMMMMMMMb          ..MMMMMMMMMM#",
			"      .MMMMMMMMM#`                  dMMMMMMMMMN      ` MMMMMMMMM\\       `    .MMMMMMMMM\\       JMMMMMMMMMMNJ.. `...dMMMMMMMMMMMM%",
			"     .MMMMMMMMM#`                   JMMMMMMMMMM-      .MMMMMMMM#             JMMMMMMMM#         WMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM#",
			"    -MMMMMMMMM@                     ,MMMMMMMMMM]      dMMMMMMMMF            .MMMMMMMMMF          UMMMMMMMMMMMMMMMMMMMD(MMMMMMMMF          `   `   `   `   `   `   `   `   `   `   `   `  `  `  `  `  `",
			"   JMMMMMMMMM@                      .MMMMMMMMMM@     .MMMMMMMMM`            .MMMMMMMMM`      `    (HMMMMMMMMMMMMMMM\"  MMMMMMMMM!",
			"  (\"\"\"\"\"\"\"\"\"\"                        7\"\"\"\"\"\"\"\"\"\"     ,\"\"\"\"\"\"\"\"^     `   `   ?\"\"\"\"\"\"\"\"^               7WMMMMMMMH9\"`   .\"\"\"\"\"\"\"\"=",
		})
	case 116:
		getWildAndTough = append(getWildAndTough, "TOUGH!!")
		drawGetWild(2, 1, []string{
			"  ggggggggggggggggggggggggggggg[                                                                                            .gggggge",
			" .MMMMMMMMMMMMMMMMMMMMMMMMMMMMM\\                                                                                            JMMMMMMt",
			" dMMMMMMMMMMMMMMMMMMMMMMMMMMMM#                                                                                            .MMMMMM#",
			" HHHHHHHHHHHMMMMMMM#HHHHHHHHHHt                                                                                            .MMMMMMF",
			"           (MMMMMMM`                    ....                                                         `                     MMMMMMM!",
			"           MMMMMMMF               ..NMMMMMMMMMMNJ,          .NNNNNNP         .NNNNNNK          ..gMMMMMMMa,  qNNNNNN`     .MMMMMMF  .gMMMMMMMNa,",
			"          .MMMMMMM\\            .uMMMMMMMMMMMMMMMMMN.        MMMMMMM%         JMMMMMMF        .MMMMMMMMMMMMMh.MMMMMMF      JMMMMMMLJMMMMMMMMMMMMMb",
			"          dMMMMMM#           .dMMMMMMMMMMMMMMMMMMMMN,   `  .MMMMMM#     `   .MMMMMMM`   `  .MMMMMMMMMMMMMMMMMMMMMMM\\     .MMMMMMMMMMMMMMMMMMMMMMM]",
			"         .MMMMMMMF          .MMMMMMMB^     .TMMMMMMMN      JMMMMMMF         -MMMMMMF      JMMMMMMM#\"`   ?HMMMMMMMM#      -MMMMMMMM@^     UMMMMMMMF",
			"    `    (MMMMMMM`     `   JMMMMMM#!         JMMMMMMM-    .MMMMMMM!         MMMMMMM\\     dMMMMMMB`        4MMMMMMM]   `  MMMMMMM#`       .MMMMMMM%     `    `    `    `    `    `    `    `    `    `",
			"      ` .MMMMMMMF         -MMMMMM@           .MMMMMMM)    -MMMMMMF     `   .MMMMMM#     JMMMMMMF          ,MMMMMMM`     .MMMMMM#         .MMMMMMM",
			"  `     .MMMMMMM\\   `   `.MMMMMMM`           .MMMMMMM!    MMMMMMM\\         dMMMMMM]    .MMMMMMM           .MMMMMMF      dMMMMMM%         JMMMMMMF",
			"     `  dMMMMMM#     `   dMMMMMMF    `  `    (MMMMMM#    .MMMMMM#         .MMMMMMM`    dMMMMMMF           dMMMMMM>     .MMMMMM#         .MMMMMMM!     `  `   `  `  `   `  `  `   `  `  `   `  `  `",
			"       .MMMMMMM]  `      MMMMMMM\\           .MMMMMMM^    dMMMMMMF    `    (MMMMMMF    .MMMMMMM)       `  .MMMMMM#      (MMMMMMF   `  `  .MMMMMM#                                                    `",
			"  `    JMMMMMMM`         MMMMMMM]  `       .MMMMMMMF    .MMMMMMM:       `.MMMMMMM>    .MMMMMMM]         .MMMMMMMt      MMMMMMM`         dMMMMMM%",
			"   `  .MMMMMMMF     `  ` MMMMMMMN.        .MMMMMMMD     .MMMMMMM,       .MMMMMMM#     .MMMMMMMM,   `  .dMMMMMMMM      .MMMMMMF         .MMMMMM#     `     `         `         `         `        `",
			"      .MMMMMMM:          -MMMMMMMMN.....gMMMMMMMM^      ,MMMMMMMMN-..JMMMMMMMMMMt      ZMMMMMMMMMNNNMMMMMMMMMMMF    ` dMMMMMM\\   `     JMMMMMMF        `     `  `      `  `      `  `      `  `",
			"    ` MMMMMMM#   `  `     ?MMMMMMMMMMMMMMMMMMMM\"        .MMMMMMMMMMMMMMMBMMMMMMM        TMMMMMMMMMMMMMM@MMMMMMM!     .MMMMMM#       ` .MMMMMMM!                                                     `",
			"     .MMMMMMMt             .YMMMMMMMMMMMMMMMY^           .WMMMMMMMMMMMD!.MMMMMMF         .TMMMMMMMMMM\"`.MMMMMMF      (MMMMMM]         .MMMMMMF",
			"      ~~~~~~~                 _\"\"\"MMMM\"\"\"=`                 ?\"\"\"\"\"\"^     ~~~~~~`              ???!     dMMMMMM\\      _~~~~~~          _~~~~~~`",
			"                                                                                    ........          .MMMMMMF",
			"                                                      `                             dMMMMMMMp       .MMMMMMM@",
			"                                                                                    ,MMMMMMMMMMNNMMMMMMMMMMD",
			"                                                  `                                  ,MMMMMMMMMMMMMMMMMMM\"",
			"                                                                    `                   7\"MMMMMMMMMM\"\"^",
		})
	default:
		getWildAndTough = []string{}
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
				"                                                                                                         `",
				"",
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
		} else {
			drawGetWild(2, 1, []string{
				"==",
				"=====",
				"=========",
				"===========",
				"==============",
				"===================",
				"==============",
				"==========",
				"========",
				"======",
			})
			drawGetWild(2, 1, []string{
				"===================",
				"================",
				"=================",
				"==============",
				"========",
				"==========",
				"==============",
				"==========",
				"========",
				"======",
			})
			drawGetWild(2, 1, []string{
				"======",
				"========",
				"==========",
				"==============",
				"==============",
				"========",
				"==========",
				"================",
				"=================",
				"===================",
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
