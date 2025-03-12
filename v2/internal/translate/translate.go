package translate

import (
	"os"

	"github.com/arturbaccarin/go-subtitle-translator/internal/subtitletranslator"
	"github.com/arturbaccarin/go-subtitle-translator/pkg/requester/nethttp"
	"github.com/arturbaccarin/go-subtitle-translator/pkg/subtitlereader/srt"
	"github.com/arturbaccarin/go-subtitle-translator/pkg/translator/deepl"
)

func SRT(filePath string, originalLang string, targetLang string) error {
	reader := srt.NewReader(filePath)
	requester := nethttp.NewNetHTTPRequester()
	apiClient := deepl.NewAPIClient(os.Getenv("DEEPL_API_KEY"), os.Getenv("DEEPL_HOSTNAME"), requester)

	subtitleTranslator := subtitletranslator.NewSubtitleTranslator(reader, apiClient, filePath)

	subtitles, err := subtitleTranslator.Translate()
	if err != nil {
		panic(err)
	}

	err = subtitleTranslator.SaveSRT(subtitles)
	if err != nil {
		panic(err)
	}

}
