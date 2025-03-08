package subtitletranslator

import (
	"fmt"
	"os"

	"github.com/arturbaccarin/go-subtitle-translator/pkg/subtitlereader"
	"github.com/arturbaccarin/go-subtitle-translator/pkg/translator"
	"github.com/arturbaccarin/go-subtitle-translator/pkg/translator/deepl/language"
	"github.com/arturbaccarin/go-subtitle-translator/pkg/translator/dto"
)

type SubtitleTranslator struct {
	fileToTranslate string
	subtitlereader  subtitlereader.SubtitleReader
	translator      translator.Translator
}

func NewSubtitleTranslator(subtitlereader subtitlereader.SubtitleReader, translator translator.Translator, fileToTranslate string) *SubtitleTranslator {
	return &SubtitleTranslator{
		subtitlereader:  subtitlereader,
		translator:      translator,
		fileToTranslate: fileToTranslate,
	}
}

func (st *SubtitleTranslator) Translate() ([]*subtitlereader.Subtitle, error) {
	subtitles, err := st.subtitlereader.Read()
	if err != nil {
		return nil, err
	}

	linesToTranslate := st.parseContent(subtitles)

	requestPayload := st.createRequestPayload(linesToTranslate)

	translatedLines, err := st.translator.Translate(requestPayload)
	if err != nil {
		return nil, err
	}

	for i, subtitle := range subtitles {
		subtitle.Content = translatedLines.Translations[i].Text
	}

	return subtitles, nil
}

func (st *SubtitleTranslator) SaveSRT(subtitles []*subtitlereader.Subtitle) error {
	filename := st.fileToTranslate + "_translated.srt"

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, subtitle := range subtitles {
		_, err := fmt.Fprintf(file, "%d\n%s\n%s\n\n",
			subtitle.Index, subtitle.Time, subtitle.Content)
		if err != nil {
			return err
		}
	}

	return nil
}

func (st *SubtitleTranslator) parseContent(subtitles []*subtitlereader.Subtitle) []string {
	linesToTranslate := make([]string, len(subtitles))

	for _, subtitle := range subtitles {
		linesToTranslate = append(linesToTranslate, subtitle.Content)
	}

	return linesToTranslate
}

func (st *SubtitleTranslator) createRequestPayload(linesToTranslate []string) dto.Request {
	return dto.Request{
		Text:       linesToTranslate,
		TargetLang: language.PT_BR_TL,
		SourceLang: language.EN_SL,
	}

}
