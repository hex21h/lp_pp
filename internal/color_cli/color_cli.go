package color_cli

import "errors"

// the function returns the color reset code in string format
func ResetColor() (colorCode string) {
	colorCode = "\033[0m"
	return
}

// the function returns the color code in string format, the following colors are supported: red, green, yellow, blue, purple, cyan, white
func GetColor(colorText string) (colorCode string, err error) {
	switch colorText {
	case "red":
		colorCode = "\033[31m"
	case "green":
		colorCode = "\033[32m"
	case "yellow":
		colorCode = "\033[33m"
	case "blue":
		colorCode = "\033[34m"
	case "purple":
		colorCode = "\033[35m"
	case "cyan":
		colorCode = "\033[36m"
	case "white":
		colorCode = "\033[37m"
	default:
		err = errors.New("incorrect argument of text color")
	}
	return
}
