package main

import (
	"os"

	"github.com/arturbaccarin/go-subtitle-translator/internal/subtitletranslator"
	"github.com/arturbaccarin/go-subtitle-translator/pkg/requester/nethttp"
	"github.com/arturbaccarin/go-subtitle-translator/pkg/subtitlereader/srt"
	"github.com/arturbaccarin/go-subtitle-translator/pkg/translator/deepl"
	"github.com/joho/godotenv"
)

func main() {
	wd, _ := os.Getwd()

	err := godotenv.Load(wd + "/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	myFile := wd + "/test.srt"

	reader := srt.NewReader(myFile)
	requester := nethttp.NewNetHTTPRequester()
	apiClient := deepl.NewAPIClient(os.Getenv("DEEPL_API_KEY"), os.Getenv("DEEPL_HOSTNAME"), requester)

	subtitleTranslator := subtitletranslator.NewSubtitleTranslator(reader, apiClient, myFile)

	subtitles, err := subtitleTranslator.Translate()
	if err != nil {
		panic(err)
	}

	err = subtitleTranslator.SaveSRT(subtitles)
	if err != nil {
		panic(err)
	}
}
