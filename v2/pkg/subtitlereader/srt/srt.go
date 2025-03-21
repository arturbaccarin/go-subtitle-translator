package srt

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/arturbaccarin/go-subtitle-translator/pkg/subtitlereader"
)

type Reader struct {
	FilePath string
}

func NewReader(filePath string) *Reader {
	return &Reader{
		FilePath: filePath,
	}
}

func (r *Reader) Read() ([]*subtitlereader.Subtitle, error) {
	log.Printf("Reading file %s", r.FilePath)

	file, err := os.Open(r.FilePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var subtitles []*subtitlereader.Subtitle
	var currentSubtitle *subtitlereader.Subtitle

	timestampRegex := regexp.MustCompile(`(\d{2}:\d{2}:\d{2},\d{3}) --> (\d{2}:\d{2}:\d{2},\d{3})`)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		index, err := strconv.Atoi(line)
		if err == nil {
			currentSubtitle = new(subtitlereader.Subtitle)
			currentSubtitle.Index = index
			subtitles = append(subtitles, currentSubtitle)
			continue
		}

		if timestampRegex.MatchString(line) {
			currentSubtitle.Time = line
			continue
		}

		if line != "" {
			currentSubtitle.Content += line + " "
		}
	}

	log.Printf("Read %d subtitles", len(subtitles))
	return subtitles, nil
}
