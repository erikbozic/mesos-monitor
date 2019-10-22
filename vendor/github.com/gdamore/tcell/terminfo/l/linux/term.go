// Generated automatically.  DO NOT HAND-EDIT.

package linux

import "github.com/gdamore/tcell/terminfo"

func init() {

	// linux console
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:         "linux",
		Colors:       8,
		Bell:         "\a",
		Clear:        "\x1b[H\x1b[J",
		ShowCursor:   "\x1b[?25h\x1b[?0c",
		HideCursor:   "\x1b[?25l\x1b[?1c",
		AttrOff:      "\x1b[m\x0f",
		Underline:    "\x1b[4m",
		Bold:         "\x1b[1m",
		Dim:          "\x1b[2m",
		Blink:        "\x1b[5m",
		Reverse:      "\x1b[7m",
		SetFg:        "\x1b[3%p1%dm",
		SetBg:        "\x1b[4%p1%dm",
		SetFgBg:      "\x1b[3%p1%d;4%p2%dm",
		PadChar:      "\x00",
		AltChars:     "++,,--..00__``aaffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}c~~",
		EnterAcs:     "\x0e",
		ExitAcs:      "\x0f",
		EnableAcs:    "\x1b)0",
		Mouse:        "\x1b[M",
		MouseMode:    "%?%p1%{1}%=%t%'h'%Pa%e%'l'%Pa%;\x1b[?1000%ga%c\x1b[?1002%ga%c\x1b[?1003%ga%c\x1b[?1006%ga%c",
		SetCursor:    "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:  "\b",
		CursorUp1:    "\x1b[A",
		KeyUp:        "\x1b[A",
		KeyDown:      "\x1b[B",
		KeyRight:     "\x1b[C",
		KeyLeft:      "\x1b[D",
		KeyInsert:    "\x1b[2~",
		KeyDelete:    "\x1b[3~",
		KeyBackspace: "\xff",
		KeyHome:      "\x1b[1~",
		KeyEnd:       "\x1b[4~",
		KeyPgUp:      "\x1b[5~",
		KeyPgDn:      "\x1b[6~",
		KeyF1:        "\x1b[[A",
		KeyF2:        "\x1b[[B",
		KeyF3:        "\x1b[[C",
		KeyF4:        "\x1b[[D",
		KeyF5:        "\x1b[[E",
		KeyF6:        "\x1b[17~",
		KeyF7:        "\x1b[18~",
		KeyF8:        "\x1b[19~",
		KeyF9:        "\x1b[20~",
		KeyF10:       "\x1b[21~",
		KeyF11:       "\x1b[23~",
		KeyF12:       "\x1b[24~",
		KeyF13:       "\x1b[25~",
		KeyF14:       "\x1b[26~",
		KeyF15:       "\x1b[28~",
		KeyF16:       "\x1b[29~",
		KeyF17:       "\x1b[31~",
		KeyF18:       "\x1b[32~",
		KeyF19:       "\x1b[33~",
		KeyF20:       "\x1b[34~",
		KeyBacktab:   "\x1b[Z",
	})
}
