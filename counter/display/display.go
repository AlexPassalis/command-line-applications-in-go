package display

type Options struct {
	showBytes   bool
	showLines   bool
	showWords   bool
	showHeaders bool
}

type NewOptionsArguments struct {
	ShowBytes   bool
	ShowWords   bool
	ShowLines   bool
	ShowHeaders bool
}

func NewOptions(arguments NewOptionsArguments) Options {
	return Options{
		showBytes:   arguments.ShowBytes,
		showWords:   arguments.ShowWords,
		showLines:   arguments.ShowLines,
		showHeaders: arguments.ShowHeaders,
	}
}

func (d Options) ShouldShowBytes() bool {
	if !d.showBytes && !d.showWords && !d.showLines && !d.showHeaders {
		return true
	}

	return d.showBytes
}

func (d Options) ShouldShowWords() bool {
	if !d.showBytes && !d.showWords && !d.showLines && !d.showHeaders {
		return true
	}

	return d.showWords
}

func (d Options) ShouldShowLines() bool {
	if !d.showBytes && !d.showWords && !d.showLines && !d.showHeaders {
		return true
	}

	return d.showLines
}

func (d Options) ShouldShowHeader() bool {
	return d.showHeaders
}
