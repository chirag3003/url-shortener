package helpers

import (
	"bytes"
	"io"
	"mime/multipart"

	"github.com/h2non/bimg"
)

func OptimiseImage(_file multipart.FileHeader) (io.Reader, error) {
	file, err := _file.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content into a byte slice
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Create a new bimg image from the byte slice
	img := bimg.NewImage(fileBytes)

	// Optimize the image by resizing and setting quality
	options := bimg.Options{
		Type:    bimg.WEBP,
		Quality: 75, // Adjust quality to reduce size
	}

	imgBytes, err := img.Process(options)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(imgBytes), nil
}
