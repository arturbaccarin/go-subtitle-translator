package front

import "net/http"

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "uploads/translated.srt")
}
