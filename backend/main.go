package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	vision "cloud.google.com/go/vision/apiv1"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    // 1. Parse the multipart form (image file)
    // 2. Upload to Firebase Storage
    // 3. Trigger Google Vision API for classification
    // 4. Return JSON response to Frontend
    fmt.Fprintf(w, "Backend is ready to process images!")
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	port := os.Getenv("PORT")
	if port == "" { port = "8080" }
	fmt.Printf("Server starting on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}