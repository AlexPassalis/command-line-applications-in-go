package display

type Options struct {
	ShowBytes  bool
	ShowWords  bool
	ShowLines  bool
	ShowHeader bool
}

func (d Options) ShouldShowBytes() bool {
	if !d.ShowBytes && !d.ShowWords && !d.ShowLines && !d.ShowHeader {
		return true
	}

	return d.ShowBytes
}

func (d Options) ShouldShowWords() bool {
	if !d.ShowBytes && !d.ShowWords && !d.ShowLines && !d.ShowHeader {
		return true
	}

	return d.ShowWords
}

func (d Options) ShouldShowLines() bool {
	if !d.ShowBytes && !d.ShowWords && !d.ShowLines && !d.ShowHeader {
		return true
	}

	return d.ShowLines
}

func (d Options) ShouldShowHeader() bool {
	return d.ShowHeader
}
