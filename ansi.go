package ansi

const (
	csi             = "\033["
	reset           = "0"
	bold            = ";1"
	italic          = ";3"
	underline       = ";4"
	blink           = ";5"
	foreground      = "3"
	foregroundLight = "9"
	background      = "4"
	backgroundLight = "10"
	black           = "0"
	red             = "1"
	green           = "2"
	yellow          = "3"
	blue            = "4"
	magenta         = "5"
	cyan            = "6"
	white           = "7"
	def             = "9"
	sgr             = "m"
)

type modifier uint

// Text modifiers.
const (
	None           = 0
	Light modifier = 1 << iota
	Bold
	Italic
	Underline
	Blink
)

var enabled = true

// Enable enables the use of ANSI escape sequences in calls to ansi functions.
func Enable() {
	enabled = true
}

// Disable disabled the use of ANSI escape sequences in calls to ansi functions.
// In short, no string manipulations will be done by the ansi functions.
func Disable() {
	enabled = false
}

// Black returns a string enclosed in the ANSI escape code for black text.
func Black(str string, mod modifier) string {
	return colorString(str, black, mod)
}

// Red returns a string enclosed in the ANSI escape code for red text.
func Red(str string, mod modifier) string {
	return colorString(str, red, mod)
}

// Green returns a string enclosed in the ANSI escape code for green text.
func Green(str string, mod modifier) string {
	return colorString(str, green, mod)
}

// Yellow returns a string enclosed in the ANSI escape code for yellow text.
func Yellow(str string, mod modifier) string {
	return colorString(str, yellow, mod)
}

// Blue returns a string enclosed in the ANSI escape code for blue text.
func Blue(str string, mod modifier) string {
	return colorString(str, blue, mod)
}

// Magenta returns a string enclosed in the ANSI escape code for magenta text.
func Magenta(str string, mod modifier) string {
	return colorString(str, magenta, mod)
}

// Cyan returns a string enclosed in the ANSI escape code for cyan text.
func Cyan(str string, mod modifier) string {
	return colorString(str, cyan, mod)
}

// White returns a string enclosed in the ANSI escape code for white text.
func White(str string, mod modifier) string {
	return colorString(str, white, mod)
}

func colorString(str, col string, mod modifier) string {
	if !enabled {
		return str
	}
	s := csi
	if mod&Light != 0 {
		s += foregroundLight
	} else {
		s += foreground
	}
	s += col
	if mod&Bold != 0 {
		s += bold
	}
	if mod&Italic != 0 {
		s += italic
	}
	if mod&Underline != 0 {
		s += underline
	}
	if mod&Blink != 0 {
		s += blink
	}
	s += sgr + str
	s += csi + reset + sgr
	return s
}
