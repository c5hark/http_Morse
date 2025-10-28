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
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func MorseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "form parsing error", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "failed to retieve the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed to read the file", http.StatusInternalServerError)
		return
	}

	originalText := string(content)

	converted, err := service.TranslateToMorse(originalText)
	if err != nil {
		http.Error(w, "failed to convert the file", http.StatusInternalServerError)
		return
	}

	fileName := time.Now().UTC().Format("20060102_150405") + filepath.Ext(handler.Filename)
	newFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "failed to create the file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString(converted)
	if err != nil {
		http.Error(w, "failed to write converted data to file", http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "Результат:\n%s\n\nИсходный текст:\n%s\n", converted, originalText)
	if err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}
