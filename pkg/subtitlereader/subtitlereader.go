package subtitlereader

type Subtitle struct {
	Index   int
	Time    string
	Content string
}

type SubtitleReader interface {
	Read() ([]*Subtitle, error)
}
