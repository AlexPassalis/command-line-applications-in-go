package display

type Options struct {
	arguments NewOptionsArguments
}

type NewOptionsArguments struct {
	ShowBytes   bool
	ShowWords   bool
	ShowLines   bool
	ShowHeaders bool
}

func NewOptions(arguments NewOptionsArguments) Options {
	return Options{
		arguments: arguments,
	}
}

func (d Options) ShouldShowBytes() bool {
	if !d.arguments.ShowBytes && !d.arguments.ShowWords && !d.arguments.ShowLines && !d.arguments.ShowHeaders {
		return true
	}

	return d.arguments.ShowBytes
}

func (d Options) ShouldShowWords() bool {
	if !d.arguments.ShowBytes && !d.arguments.ShowWords && !d.arguments.ShowLines && !d.arguments.ShowHeaders {
		return true
	}

	return d.arguments.ShowWords
}

func (d Options) ShouldShowLines() bool {
	if !d.arguments.ShowBytes && !d.arguments.ShowWords && !d.arguments.ShowLines && !d.arguments.ShowHeaders {
		return true
	}

	return d.arguments.ShowLines
}

func (d Options) ShouldShowHeader() bool {
	return d.arguments.ShowHeaders
}
