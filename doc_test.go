package ansi

import "fmt"

func ExampleBlack() {
	fmt.Println(Black("This text will be black", None))
}

func ExampleBlack_light() {
	fmt.Println(Black("This text will be light black", Light))
}

func ExampleBlack_bold() {
	fmt.Println(Black("This text will be bold and black", Bold))
}

func ExampleBlack_italic() {
	fmt.Println(Black("This text will be italic and black", Italic))
}

func ExampleBlack_underline() {
	fmt.Println(Black("This text will be underlined and black", Underline))
}

func ExampleBlack_blink() {
	fmt.Println(Black("This text will be blinking and black", Blink))
}

func ExampleBlack_combined() {
	fmt.Println(Black("This text will be bold, underlined and black", Bold|Underline))
}
