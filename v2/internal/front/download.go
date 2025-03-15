package front

import (
	"net/http"
	"os"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	wd, _ := os.Getwd()
	name := wd + "/uploads/translated.srt"

	w.Header().Set("Content-Disposition", "attachment; filename=translated.srt")
	http.ServeFile(w, r, name)
}
