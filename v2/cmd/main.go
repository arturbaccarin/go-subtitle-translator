package main

import (
	"log"
	"net/http"
	"os"

	"github.com/arturbaccarin/go-subtitle-translator/internal/front"
	"github.com/joho/godotenv"
)

func main() {
	wd, _ := os.Getwd()

	err := godotenv.Load(wd + "/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", front.HomepageHandler)
	http.HandleFunc("/translate", front.TranslateHandler)
	http.HandleFunc("/download", front.DownloadHandler)

	log.Println("Server started on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// func main() {
// 	wd, _ := os.Getwd()

// 	err := godotenv.Load(wd + "/.env")
// 	if err != nil {
// 		panic("Error loading .env file")
// 	}

// 	myFile := wd + "/test.srt"

// 	reader := srt.NewReader(myFile)
// 	requester := nethttp.NewNetHTTPRequester()
// 	apiClient := deepl.NewAPIClient(os.Getenv("DEEPL_API_KEY"), os.Getenv("DEEPL_HOSTNAME"), requester)

// 	subtitleTranslator := subtitletranslator.NewSubtitleTranslator(reader, apiClient, myFile)

// 	subtitles, err := subtitleTranslator.Translate()
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = subtitleTranslator.SaveSRT(subtitles)
// 	if err != nil {
// 		panic(err)
// 	}
// }
