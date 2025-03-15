package translate

import (
	"fmt"
	"os"

	"github.com/arturbaccarin/go-subtitle-translator/internal/subtitletranslator"
	"github.com/arturbaccarin/go-subtitle-translator/pkg/requester/nethttp"
	"github.com/arturbaccarin/go-subtitle-translator/pkg/subtitlereader/srt"
	"github.com/arturbaccarin/go-subtitle-translator/pkg/translator/deepl"
)

func SRT(originalFilePath string, translatedFilePath string, originalLang string, targetLang string) error {
	reader := srt.NewReader(originalFilePath)
	requester := nethttp.NewNetHTTPRequester()
	apiClient := deepl.NewAPIClient(os.Getenv("DEEPL_API_KEY"), os.Getenv("DEEPL_HOSTNAME"), requester)

	subtitleTranslator := subtitletranslator.NewSubtitleTranslator(reader, apiClient, originalFilePath)

	subtitles, err := subtitleTranslator.Translate()
	if err != nil {
		return fmt.Errorf("error translating subtitles: %v", err)
	}

	err = subtitleTranslator.SaveSRT(subtitles, translatedFilePath)
	if err != nil {
		panic(err)
	}

	return nil
}
