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

	// ColorEsc start color
	ColorEsc = "\033["
	// ColorClear clear color
	ColorClear = ColorEsc + "0m"
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
	// ColorWhite white
	ColorWhite = ColorEsc + "37m"
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
)
