package sakescript

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func ZipIndexer(path string) ([]IndexEntry, error) {
	var index []IndexEntry
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".zip") {
			manifest, err := extractFile(path, "manifest.json")
			if err != nil {
				return err
			}

			if manifest != nil {
				hash, err := computeHash(path)
				if err != nil {
					return err
				}
				index = append(index, IndexEntry{Path: path, Sha256: hash, Manifest: *manifest})
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return index, nil
}

// Sort sort index by created date desc
func SortIndex(index []IndexEntry) {
	sort.Slice(index, func(i, j int) bool {
		return index[i].Manifest.Created > index[j].Manifest.Created
	})
}

func WriteIndex(index []IndexEntry, output string) error {
	indexData, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(output, indexData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func extractFile(zipPath, filename string) (*Manifest, error) {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == filename {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()

			manifestData, err := io.ReadAll(rc)
			if err != nil {
				return nil, err
			}

			var manifest Manifest
			if err := json.Unmarshal(manifestData, &manifest); err != nil {
				return nil, err
			}
			return &manifest, nil
		}
	}
	return nil, fmt.Errorf("file %s not found in %s", filename, zipPath)
}

func computeHash(zipPath string) (string, error) {
	file, err := os.Open(zipPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
