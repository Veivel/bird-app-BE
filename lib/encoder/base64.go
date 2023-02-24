package encoder

import (
	"encoding/base64"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func GetBase64(file multipart.File) (base64Encoding string) {
	// Read the entire file into a byte slice
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	return base64Encoding
}
