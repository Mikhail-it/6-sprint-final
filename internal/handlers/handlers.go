package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("..\\index.html")
	if err != nil {
		http.Error(w, "Ошибка чтения index.html", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(30)

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	converted := service.Convert(string(data))

	fileName := fmt.Sprintf("%s%s", time.Now().UTC().Format("2006-01-02_15-04-05"), filepath.Ext("output.txt"))
	filePath := filepath.Join("..", fileName)

	err = os.WriteFile(filePath, []byte(converted), 0755)
	if err != nil {
		http.Error(w, "Ошибка при записи в файл", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(converted))
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/upload", UploadHandler)
}
