package logging

const (
	// Empty empty string
	Empty string = ""
	// NewLine line terminator
	NewLine string = "\n"
	// OneMega 1M
	OneMega int64 = 1024 * 1024
	// DefaultLogFileSize defautl log file size
	DefaultLogFileSize int64 = OneMega * 2

	// LeftBracket left bracket
	LeftBracket = "["
	// RightBracket right bracket
	RightBracket = "]"
	// Space space
	Space = " "

	// StyleSeparator sign uses to separator styles
	StyleSeparator = ","

	// Name name
	Name = "Name"
	// Time time
	Time = "Time"
	// Level level
	Level = "Level"
	// Indent indent
	Indent = "Indent"
	// Message message
	Message = "Message"
	// Compose compose
	Compose = "Compose"

	// NameLayout name layout
	NameLayout = "NameLayout"
	// TimeLayout time layout
	TimeLayout = "TimeLayout"
	// LevelLayout level
	LevelLayout = "LevelLayout"
	// IndentLayout indent
	IndentLayout = "IndentLayout"
	// MessageLayout message
	MessageLayout = "MessageLayout"
	// ComposeLayout compose
	ComposeLayout = "ComposeLayout"

	/*====================================== Colors =============================*/

	// StyleEsc start color
	StyleEsc = "\033["
	// StyleEnd style end indicator, a style could be 37m, or combined style 1;2;32m
	StyleEnd = "m"
	// ColorDefaultText default text color
	ColorDefaultText = "39"
	// ColorBlack black
	ColorBlack = "30"
	// ColorRed red
	ColorRed = "31"
	// ColorGreen green
	ColorGreen = "32"
	// ColorYellow yellow
	ColorYellow = "33"
	// ColorBlue blue
	ColorBlue = "34"
	// ColorMagenta magenta
	ColorMagenta = "35"
	// ColorCyan cyan
	ColorCyan = "36"
	// ColorLightGray light gray
	ColorLightGray = "37"
	// ColorBgDefault default bg
	ColorBgDefault = "49"
	// ColorBgBlack black
	ColorBgBlack = "40"
	// ColorBgRed red
	ColorBgRed = "41"
	// ColorBgGreen green
	ColorBgGreen = "42"
	// ColorBgYellow yellow
	ColorBgYellow = "43"
	// ColorBgBlue blue
	ColorBgBlue = "44"
	// ColorBgMagenta magenta
	ColorBgMagenta = "45"
	// ColorBgCyan cyan
	ColorBgCyan = "46"
	// ColorBgLightGray light gray
	ColorBgLightGray = "47"
	// ColorBgDarkGray dark gray
	ColorBgDarkGray = "100"
	// ColorBgLightRed light red
	ColorBgLightRed = "101"
	// ColorBgLightGreen light green
	ColorBgLightGreen = "102"
	// ColorBgLightYellow light yellow
	ColorBgLightYellow = "103"
	// ColorBgLightBlue light blue
	ColorBgLightBlue = "104"
	// ColorBgLightMagenta light magenta
	ColorBgLightMagenta = "105"
	// ColorBgLightCyan light cyan
	ColorBgLightCyan = "106"
	// ColorBgWhite bg white
	ColorBgWhite = "107"
	// ColorDarkGray dark gray
	ColorDarkGray = "90"
	// ColorLightRed light red
	ColorLightRed = "91"
	// ColorLightGreen light green
	ColorLightGreen = "92"
	// ColorLightYellow light yellow
	ColorLightYellow = "93"
	// ColorLightBlue lightblue
	ColorLightBlue = "94"
	// ColorLightMagenta light magenta
	ColorLightMagenta = "95"
	// ColorLightCyan light cyan
	ColorLightCyan = "96"
	// ColorWhite white
	ColorWhite = "97"
	// ColorBrightBlack bright black
	ColorBrightBlack = "30;1"
	// ColorBrightRed bright red
	ColorBrightRed = "31;1"
	// ColorBrightGreen bright green
	ColorBrightGreen = "32;1"
	// ColorBrightYellow brightyellow
	ColorBrightYellow = "33;1"
	// ColorBrightBlue bright blue
	ColorBrightBlue = "34;1"
	// ColorBrightMagenta bright magenta
	ColorBrightMagenta = "35;1"
	// ColorBrightCyan bright cyan
	ColorBrightCyan = "36;1"
	// ColorBrightWhite bright white
	ColorBrightWhite = "37;1"
	// Color256Start 256 color start
	Color256Start = "38;5;"
	// Color256End 256 color end
	Color256End = ""
	// Color256BgStart 256 color start
	Color256BgStart = "48;5;"
	// Color256BgEnd 256 color end
	Color256BgEnd = ""

	/*====================================== Styles =============================*/

	// StyleBold bold style
	StyleBold = "1"
	// StyleDim dim
	StyleDim = "2"
	// StyleUnderline underline
	StyleUnderline = "4"
	// StyleBlink blink
	StyleBlink = "5"
	// StyleReverse reverse
	StyleReverse = "7"
	// StyleHidden hidden
	StyleHidden = "8"

	// ResetAllStyle clear color
	ResetAllStyle = StyleEsc + "0m"
	// ResetBold reset bold
	ResetBold = StyleEsc + "21m"
	// ResetDim reset dim
	ResetDim = StyleEsc + "22m"
	// ResetUnderline reset underline
	ResetUnderline = StyleEsc + "24m"
	// ResetBlink reset blink
	ResetBlink = StyleEsc + "25m"
	// ResetReverse reset reverse
	ResetReverse = StyleEsc + "27m"
	// ResetHidden reset hidden
	ResetHidden = StyleEsc + "28m"

	// ForegroundPrefix foreground fix
	ForegroundPrefix = "Fore"
	// BackgroundPrefix background prefix
	BackgroundPrefix = "Bg"
)

var (
	// KnownStyles know styles
	KnownStyles = map[string]string{
		"Bold":      "1",
		"Dim":       "2",
		"Underline": "4",
		"Blink":     "5",
		"Reverse":   "6",
		"Hidden":    "8",
	}
	// KnownColors known colors
	KnownColors = map[string]string{
		"DefaultText":      "39",
		"ForeBlack":        "30",
		"ForeRed":          "31",
		"ForeGreen":        "32",
		"ForeYellow":       "33",
		"ForeBlue":         "34",
		"ForeMagenta":      "35",
		"ForeCyan":         "36",
		"ForeLightGray":    "37",
		"ForeDarkGray":     "90",
		"ForeLightRed":     "91",
		"ForeLightGreen":   "92",
		"ForeLightYellow":  "93",
		"ForeLightBlue":    "94",
		"ForeLightMagenta": "95",
		"ForeLightCyan":    "96",
		"ForeLightWhite":   "97",
		"BgDefault":        "49",
		"BgBlack":          "40",
		"BgRed":            "41",
		"BgGreen":          "42",
		"BgYellow":         "43",
		"BgBlue":           "44",
		"BgMagenta":        "45",
		"BgCyan":           "46",
		"BgLightGray":      "47",
		"BgDarkGray":       "100",
		"BgLightRed":       "101",
		"BgLightGreen":     "102",
		"BgLightYellow":    "103",
		"BgLightBlue":      "104",
		"BgLightMagenta":   "105",
		"BgLightCyan":      "106",
		"BgLightWhite":     "107",
		"BrightBlack":      "30;1",
		"BrightRed":        "31;1",
		"BrightGreen":      "32;1",
		"BrightYellow":     "33;1",
		"BrightBlue":       "34;1",
		"BrightMagenta":    "35;1",
		"BrightCyan":       "36;1",
		"BrightWhite":      "37;1",
	}
)
