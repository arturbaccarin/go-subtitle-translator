package front

import (
	"net/http"
	"os"

	"github.com/arturbaccarin/go-subtitle-translator/internal/translate"
)

func translateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(500 * 1024) // 500 KB limit
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		originalLang := r.FormValue("originalLang")
		targetLang := r.FormValue("targetLang")

		file, _, err := r.FormFile("srtFile")
		if err != nil {
			http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		originalFilePath := "uploads/original.srt"
		dst, err := os.Create(originalFilePath)
		if err != nil {
			http.Error(w, "Unable to save file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.ReadFrom(file)
		if err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}

		translatedFilePath := "uploads/translated.srt"
		err = translate.SRT(originalFilePath, translatedFilePath, originalLang, targetLang)
		if err != nil {
			http.Error(w, "Translation failed", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/download", http.StatusSeeOther)
	}
}
