package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "/var/www/html/upload.html")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := header.Filename
	dstPath := filepath.Join("/var/www/html/uploads", filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	// Tell the browser we're sending HTML
	w.Header().Set("Content-Type", "text/html")

	// Output clickable link
	fmt.Fprintf(w, "File uploaded! Access it at <a href=\"/uploads/%s\">/uploads/%s</a>", filename, filename)
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/upload", http.StatusSeeOther)
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
