package sava

import (
	"fmt"
	"net/http"

	"ascii-art-web/asciiArt"
)

func HandleasciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	m, err := asciiArt.LoadBannerMap("bannerfiles/" + banner + ".txt")
	if err != nil {
		// return error 500
		errorMessage := fmt.Sprintf("Internal Server Error: \n%v", err)
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}
	art, err := asciiArt.PrintLineBanner(text, m)
	if err != nil {
		// return error 500// return error 500
		errorMessage := fmt.Sprintf("Internal Server Error: \n%v", err)
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}

	data := struct {
		AsciiArt string
	}{
		AsciiArt: art,
	}

	t.ExecuteTemplate(w, "template.html", data)
}
