package ansi

import (
	"testing"

	"github.com/zhevron/match"
)

func TestBlack(t *testing.T) {
	str := Black("test", None)
	match.GreaterThan(t, len(str), 0)
	match.Equals(t, str, "\033[30mtest\033[0m")
}

func TestRed(t *testing.T) {
	str := Red("test", None)
	match.GreaterThan(t, len(str), 0)
	match.Equals(t, str, "\033[31mtest\033[0m")
}

func TestGreen(t *testing.T) {
	str := Green("test", None)
	match.GreaterThan(t, len(str), 0)
	match.Equals(t, str, "\033[32mtest\033[0m")
}

func TestYellow(t *testing.T) {
	str := Yellow("test", None)
	match.GreaterThan(t, len(str), 0)
	match.Equals(t, str, "\033[33mtest\033[0m")
}

func TestBlue(t *testing.T) {
	str := Blue("test", None)
	match.GreaterThan(t, len(str), 0)
	match.Equals(t, str, "\033[34mtest\033[0m")
}

func TestMagenta(t *testing.T) {
	str := Magenta("test", None)
	match.GreaterThan(t, len(str), 0)
	match.Equals(t, str, "\033[35mtest\033[0m")
}

func TestCyan(t *testing.T) {
	str := Cyan("test", None)
	match.GreaterThan(t, len(str), 0)
	match.Equals(t, str, "\033[36mtest\033[0m")
}

func TestWhite(t *testing.T) {
	str := White("test", None)
	match.GreaterThan(t, len(str), 0)
	match.Equals(t, str, "\033[37mtest\033[0m")
}

func TestColorString_Light(t *testing.T) {
	str := Black("test", Light)
	match.GreaterThan(t, len(str), 0)
	match.Equals(t, str, "\033[90mtest\033[0m")
}

func TestColorString_Bold(t *testing.T) {
	str := Black("test", Bold)
	match.GreaterThan(t, len(str), 0)
	match.Equals(t, str, "\033[30;1mtest\033[0m")
}

func TestColorString_Italic(t *testing.T) {
	str := Black("test", Italic)
	match.GreaterThan(t, len(str), 0)
	match.Equals(t, str, "\033[30;3mtest\033[0m")
}

func TestColorString_Underline(t *testing.T) {
	str := Black("test", Underline)
	match.GreaterThan(t, len(str), 0)
	match.Equals(t, str, "\033[30;4mtest\033[0m")
}

func TestColorString_Blink(t *testing.T) {
	str := Black("test", Blink)
	match.GreaterThan(t, len(str), 0)
	match.Equals(t, str, "\033[30;5mtest\033[0m")
}

func TestColorString_Combined(t *testing.T) {
	str := Black("test", Bold|Underline)
	match.GreaterThan(t, len(str), 0)
	match.Equals(t, str, "\033[30;1;4mtest\033[0m")
}

func TestColorString_Disabled(t *testing.T) {
	Disable()
	str := Black("test", None)
	match.GreaterThan(t, len(str), 0)
	match.Equals(t, str, "test")
	Enable()
}
