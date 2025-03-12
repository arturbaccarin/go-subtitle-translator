package front

import (
	"net/http"
	"os"

	"github.com/arturbaccarin/go-subtitle-translator/internal/translate"
)

func translateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse the form data (including file upload)
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

		uploadedFilePath := "uploads/uploaded.srt"
		dst, err := os.Create(uploadedFilePath)
		if err != nil {
			http.Error(w, "Unable to save file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Copy the uploaded file content into the created file
		_, err = dst.ReadFrom(file)
		if err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}

		// Simulate translation: Translate the file (you can replace this with actual translation logic)
		translatedSrtPath := "uploads/translated.srt"
		err = translate.SRT(uploadedFilePath, translatedSrtPath, originalLang, targetLang)
		if err != nil {
			http.Error(w, "Translation failed", http.StatusInternalServerError)
			return
		}

		// Provide the download link for the translated file
		http.Redirect(w, r, "/download", http.StatusSeeOther)
	}
}
