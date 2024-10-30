package sakescript

import (
	"archive/zip"
	"encoding/json"
	"fmt"
)

func Zip(zipWriter *zip.Writer, story *Story, manifest *Manifest) error {
	storyJSON, err := json.Marshal(story)
	if err != nil {
		return err
	}

	manifestJSON, err := json.Marshal(manifest)
	if err != nil {
		return err
	}

	if err = addFileToZip(zipWriter, MainFile, storyJSON); err != nil {
		return err
	}

	if err := addFileToZip(zipWriter, ManifestFile, manifestJSON); err != nil {
		return err
	}

	if err := zipWriter.Close(); err != nil {
		return err
	}

	return nil
}

// addFileToZip adds a file with the given name and contents to the ZIP archive
func addFileToZip(zipWriter *zip.Writer, filename string, content []byte) error {
	fileWriter, err := zipWriter.Create(filename)
	if err != nil {
		return fmt.Errorf("create %s in zip file failed: %v", filename, err)
	}

	_, err = fileWriter.Write(content)
	if err != nil {
		return fmt.Errorf("write %s to zip file failed: %v", filename, err)
	}

	return nil
}
