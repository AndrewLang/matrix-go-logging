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

	// ColorEsc start color
	ColorEsc = "\033["
	// ColorDefaultText default text color
	ColorDefaultText = ColorEsc + "39m"
	// ColorBlack black
	ColorBlack = ColorEsc + "30m"
	// ColorRed red
	ColorRed = ColorEsc + "31m"
	// ColorGreen green
	ColorGreen = ColorEsc + "32m"
	// ColorYellow yellow
	ColorYellow = ColorEsc + "33m"
	// ColorBlue blue
	ColorBlue = ColorEsc + "34m"
	// ColorMagenta magenta
	ColorMagenta = ColorEsc + "35m"
	// ColorCyan cyan
	ColorCyan = ColorEsc + "36m"
	// ColorLightGray light gray
	ColorLightGray = ColorEsc + "37m"
	// ColorBgDefault default bg
	ColorBgDefault = ColorEsc + "49m"
	// ColorBgBlack black
	ColorBgBlack = ColorEsc + "40m"
	// ColorBgRed red
	ColorBgRed = ColorEsc + "41m"
	// ColorBgGreen green
	ColorBgGreen = ColorEsc + "42m"
	// ColorBgYellow yellow
	ColorBgYellow = ColorEsc + "43m"
	// ColorBgBlue blue
	ColorBgBlue = ColorEsc + "44m"
	// ColorBgMagenta magenta
	ColorBgMagenta = ColorEsc + "45m"
	// ColorBgCyan cyan
	ColorBgCyan = ColorEsc + "46m"
	// ColorBgLightGray light gray
	ColorBgLightGray = ColorEsc + "47m"
	// ColorBgDarkGray dark gray
	ColorBgDarkGray = ColorEsc + "100m"
	// ColorBgLightRed light red
	ColorBgLightRed = ColorEsc + "101m"
	// ColorBgLightGreen light green
	ColorBgLightGreen = ColorEsc + "102m"
	// ColorBgLightYellow light yellow
	ColorBgLightYellow = ColorEsc + "103m"
	// ColorBgLightBlue light blue
	ColorBgLightBlue = ColorEsc + "104m"
	// ColorBgLightMagenta light magenta
	ColorBgLightMagenta = ColorEsc + "105m"
	// ColorBgLightCyan light cyan
	ColorBgLightCyan = ColorEsc + "106m"
	// ColorBgWhite bg white
	ColorBgWhite = ColorEsc + "107m"
	// ColorDarkGray dark gray
	ColorDarkGray = ColorEsc + "90m"
	// ColorLightRed light red
	ColorLightRed = ColorEsc + "91m"
	// ColorLightGreen light green
	ColorLightGreen = ColorEsc + "92m"
	// ColorLightYellow light yellow
	ColorLightYellow = ColorEsc + "93m"
	// ColorLightBlue lightblue
	ColorLightBlue = ColorEsc + "94m"
	// ColorLightMagenta light magenta
	ColorLightMagenta = ColorEsc + "95m"
	// ColorLightCyan light cyan
	ColorLightCyan = ColorEsc + "96m"
	// ColorWhite white
	ColorWhite = ColorEsc + "97m"
	// ColorBrightBlack bright black
	ColorBrightBlack = ColorEsc + "30;1m"
	// ColorBrightRed bright red
	ColorBrightRed = ColorEsc + "31;1m"
	// ColorBrightGreen bright green
	ColorBrightGreen = ColorEsc + "32;1m"
	// ColorBrightYellow brightyellow
	ColorBrightYellow = ColorEsc + "33;1m"
	// ColorBrightBlue bright blue
	ColorBrightBlue = ColorEsc + "34;1m"
	// ColorBrightMagenta bright magenta
	ColorBrightMagenta = ColorEsc + "35;1m"
	// ColorBrightCyan bright cyan
	ColorBrightCyan = ColorEsc + "36;1m"
	// ColorBrightWhite bright white
	ColorBrightWhite = ColorEsc + "37;1m"
	// Color256Start 256 color start
	Color256Start = ColorEsc + "38;5;"
	// Color256End 256 color end
	Color256End = "m"
	// Color256BgStart 256 color start
	Color256BgStart = ColorEsc + "48;5;"
	// Color256BgEnd 256 color end
	Color256BgEnd = "m"

	/*====================================== Styles =============================*/

	// StyleBold bold style
	StyleBold = ColorEsc + "1m"
	// StyleDim dim
	StyleDim = ColorEsc + "2m"
	// StyleUnderline underline
	StyleUnderline = ColorEsc + "4m"
	// StyleBlink blink
	StyleBlink = ColorEsc + "5m"
	// StyleReverse reverse
	StyleReverse = ColorEsc + "7m"
	// StyleHidden hidden
	StyleHidden = ColorEsc + "8m"

	// ResetAllStyle clear color
	ResetAllStyle = ColorEsc + "0m"
	// ResetBold reset bold
	ResetBold = ColorEsc + "21m"
	// ResetDim reset dim
	ResetDim = ColorEsc + "22m"
	// ResetUnderline reset underline
	ResetUnderline = ColorEsc + "24m"
	// ResetBlink reset blink
	ResetBlink = ColorEsc + "25m"
	// ResetReverse reset reverse
	ResetReverse = ColorEsc + "27m"
	// ResetHidden reset hidden
	ResetHidden = ColorEsc + "28m"
)
