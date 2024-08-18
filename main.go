package coteco

var (
	UseColors = true // Control color usage
)

func Reset() string {
	return colorize("\033[0m")
}

/////////////
// Special //
/////////////

func Bold() string {
	return colorize("\033[1m")
}

func Underline() string {
	return colorize("\033[4m")
}

/////////////////
// Text colors //
/////////////////

func Black() string {
	return colorize("\033[30m")
}

func Red() string {
	return colorize("\033[0;31m")
}

func Green() string {
	return colorize("\033[0;32m")
}

func Yellow() string {
	return colorize("\033[38;5;228m")
}

func Orange() string {
	return colorize("\033[38;5;166m")
}

func Blue() string {
	return colorize("\033[34m")
}

func Magenta() string {
	return colorize("\033[35m")
}

func Cyan() string {
	return colorize("\033[36m")
}

func GreenCyan50() string {
	return colorize("\033[38;5;50m")
}

func GreenCyan49() string {
	return colorize("\033[38;5;49m")
}

func Gray254() string {
	return colorize("\033[38;5;254m")
}

func Gray() string {
	return colorize("\033[38;5;253m")
}

func Gray250() string {
	return colorize("\033[38;5;250m")
}

func Gray247() string {
	return colorize("\033[38;5;247m")
}

func White() string {
	return colorize("\033[97m")
}

///////////////////////
// Background colors //
///////////////////////

func BlackBackground() string {
	return colorize("\033[40m")
}

func RedBackground() string {
	return colorize("\033[41m")
}

func GreenBackground() string {
	return colorize("\033[42m")
}

func YellowBackground() string {
	return colorize("\033[43m")
}

func BlueBackground() string {
	return colorize("\033[44m")
}

func MagentaBackground() string {
	return colorize("\033[45m")
}

func CyanBackground() string {
	return colorize("\033[46m")
}

func GrayBackground() string {
	return colorize("\033[47m")
}

func WhiteBackground() string {
	return colorize("\033[107m")
}

///////////////////////

func colorize(colorCode string) string {
	if UseColors {
		return colorCode
	}
	return ""
}

func With(color func() string, text string) string {
	return color() + text + Reset()
}
func WithCyan(text string) string {
	return With(Cyan, text)
}
func WithGray(text string) string {
	return With(Gray, text)
}
func WithGray250(text string) string {
	return With(Gray250, text)
}
func WithGray247(text string) string {
	return With(Gray247, text)
}
func WithGreenCyan49(text string) string {
	return With(GreenCyan49, text)
}
func WithGreen(text string) string {
	return With(Green, text)
}
func WithMagenta(text string) string {
	return With(Magenta, text)
}
func WithBlue(text string) string {
	return With(Blue, text)
}
func WithRed(text string) string {
	return With(Red, text)
}
func WithYellow(text string) string {
	return With(Yellow, text)
}
func WithOrange(text string) string {
	return With(Orange, text)
}
