package front

import (
	"net/http"
	"path/filepath"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	filename := filepath.Base(translatedFilePath)

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	http.ServeFile(w, r, translatedFilePath)
}
