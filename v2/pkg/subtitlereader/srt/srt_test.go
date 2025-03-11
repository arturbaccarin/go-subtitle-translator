package srt

import (
	"os"
	"testing"

	"github.com/arturbaccarin/go-subtitle-translator/pkg/subtitlereader"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		wantErr  bool
		wantSubs []*subtitlereader.Subtitle
	}{
		{
			name:     "valid SRT file",
			filePath: "testdata/valid.srt",
			wantErr:  false,
			wantSubs: []*subtitlereader.Subtitle{
				{Index: 1, Time: "00:00:01,000 --> 00:00:04,000", Content: "Hello world! "},
				{Index: 2, Time: "00:00:05,000 --> 00:00:08,000", Content: "This is a test. "},
			},
		},
		{
			name:     "invalid file path",
			filePath: "testdata/invalid.srt",
			wantErr:  true,
		},
		{
			name:     "empty SRT file",
			filePath: "testdata/empty.srt",
			wantErr:  false,
			wantSubs: []*subtitlereader.Subtitle{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewReader(tt.filePath)
			gotSubs, err := r.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !compareSubtitles(gotSubs, tt.wantSubs) {
				t.Errorf("Reader.Read() = %v, want %v", gotSubs, tt.wantSubs)
			}
		})
	}
}

func compareSubtitles(got, want []*subtitlereader.Subtitle) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i].Index != want[i].Index || got[i].Time != want[i].Time || got[i].Content != want[i].Content {
			return false
		}
	}
	return true
}

func TestMain(m *testing.M) {
	// Setup test data
	os.MkdirAll("testdata", 0755)
	os.WriteFile("testdata/valid.srt", []byte(`1
00:00:01,000 --> 00:00:04,000
Hello world!

2
00:00:05,000 --> 00:00:08,000
This is a test.
`), 0644)
	os.WriteFile("testdata/empty.srt", []byte(``), 0644)

	// Run tests
	code := m.Run()

	// Cleanup test data
	os.RemoveAll("testdata")

	os.Exit(code)
}
