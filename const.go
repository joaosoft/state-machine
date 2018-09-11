package color

const Escape = "\x1b"

type Format int
type Foreground int
type Background int

const (
	// format
	FormatReset Format = iota
	FormatBold
	FormatFaint
	FormatItalic
	FormatUnderline
	FormatBlinkSlow
	FormatBlinkRapid
	FormatReverseVideo
	FormatConcealed
	FormatCrossedOut
)

const (
	// foreground text colors
	ForegroundBlack Foreground = iota + 30
	ForegroundRed
	ForegroundGreen
	ForegroundYellow
	ForegroundBlue
	ForegroundMagenta
	ForegroundCyan
	ForegroundWhite
)

const (
	// foreground hi-intensity text colors
	ForegroundHiBlack Foreground = iota + 90
	ForegroundHiRed
	ForegroundHiGreen
	ForegroundHiYellow
	ForegroundHiBlue
	ForegroundHiMagenta
	ForegroundHiCyan
	ForegroundHiWhite
)

const (
	// background text colors
	BackgroundBlack Background = iota + 40
	BackgroundRed
	BackgroundGreen
	BackgroundYellow
	BackgroundBlue
	BackgroundMagenta
	BackgroundCyan
	BackgroundWhite
)

const (
	// background hi-intensity text colors
	BackgroundHiBlack Background = iota + 100
	BackgroundHiRed
	BackgroundHiGreen
	BackgroundHiYellow
	BackgroundHiBlue
	BackgroundHiMagenta
	BackgroundHiCyan
	BackgroundHiWhite
)
