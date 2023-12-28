package sakescript

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Zip(dest string, story *Story, manifest *Manifest) error {
	// Marshal the Story and Manifest into JSON
	storyJSON, err := json.Marshal(story)
	if err != nil {
		return err
	}
	manifestJSON, err := json.Marshal(manifest)
	if err != nil {
		return err
	}

	// Create a buffer to hold the ZIP file contents
	buf := new(bytes.Buffer)

	// Create a new ZIP archive in the buffer
	zipWriter := zip.NewWriter(buf)

	// Add main.json to the ZIP
	addFileToZip(zipWriter, "main.json", storyJSON)

	// Add manifest.json to the ZIP
	addFileToZip(zipWriter, "manifest.json", manifestJSON)

	// Close the ZIP writer to flush the contents to the buffer
	if err := zipWriter.Close(); err != nil {
		return err
	}

	// Write the buffer contents to a file
	file, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = buf.WriteTo(file)
	if err != nil {
		return err
	}

	return nil
}

// addFileToZip adds a file with the given name and contents to the ZIP archive
func addFileToZip(zipWriter *zip.Writer, filename string, content []byte) {
	fileWriter, err := zipWriter.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create entry %s in zip file: %v", filename, err)
	}
	_, err = fileWriter.Write(content)
	if err != nil {
		log.Fatalf("Failed to write content to %s in zip file: %v", filename, err)
	}
}

func GetIDs(path, suffix string) ([]string, error) {
	var ids []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, suffix) {
			ids = append(ids, strings.TrimSuffix(info.Name(), suffix))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return ids, nil
}
