package handlers

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "cannot parse form", http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "cannot get file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "cannot read file", http.StatusInternalServerError)
		return
	}

	converted, err := service.ConvertAutoDetect(string(data))
	if err != nil {
		http.Error(w, "conversion failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	filename := time.Now().UTC().Format("20060102_150405") + ".txt"
	f, err := os.Create(filename)
	if err != nil {
		http.Error(w, "cannot create result file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(converted); err != nil {
		http.Error(w, "cannot write to file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(converted))
}
