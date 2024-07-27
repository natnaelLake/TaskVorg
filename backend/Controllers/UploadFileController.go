package Controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form to retrieve file data
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Retrieve the file from the form
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create upload directory if it doesn't exist
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		http.Error(w, "Failed to create upload directory", http.StatusInternalServerError)
		return
	}

	// Create a new file with the original name
	fileName := fileHeader.Filename
	filePath := filepath.Join(uploadDir, fileName)
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Copy the file contents to the new file
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	// Respond with the file URL or identifier
	fileURL := fmt.Sprintf(fileName)
	fmt.Fprintf(w, `{"url": "%s"}`, fileURL)
}
