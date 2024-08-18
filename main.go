package coteco

var (
	useColors = true // Control color usage

	Reset = colorize("\033[0m")

	/////////////
	// Special //
	/////////////

	Bold      = colorize("\033[1m")
	Underline = colorize("\033[4m")

	/////////////////
	// Text colors //
	/////////////////

	Black       = colorize("\033[30m")
	Red         = colorize("\033[0;31m")
	Green       = colorize("\033[0;32m")
	Yellow      = colorize("\033[38;5;228m")
	Orange      = colorize("\033[38;5;166m")
	Blue        = colorize("\033[34m")
	Magenta     = colorize("\033[35m")
	Cyan        = colorize("\033[36m")
	GreenCyan50 = colorize("\033[38;5;50m")
	GreenCyan49 = colorize("\033[38;5;49m")
	Gray254     = colorize("\033[38;5;254m")
	Gray        = colorize("\033[38;5;253m")
	Gray250     = colorize("\033[38;5;250m")
	Gray247     = colorize("\033[38;5;247m")
	White       = colorize("\033[97m")

	///////////////////////
	// Background colors //
	///////////////////////

	BlackBackground   = colorize("\033[40m")
	RedBackground     = colorize("\033[41m")
	GreenBackground   = colorize("\033[42m")
	YellowBackground  = colorize("\033[43m")
	BlueBackground    = colorize("\033[44m")
	MagentaBackground = colorize("\033[45m")
	CyanBackground    = colorize("\033[46m")
	GrayBackground    = colorize("\033[47m")
	WhiteBackground   = colorize("\033[107m")
)

func colorize(colorCode string) string {
	if useColors {
		return colorCode
	}
	return ""
}

func With(color, text string) string {
	return color + text + Reset
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
