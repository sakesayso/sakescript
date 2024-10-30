package sakescript

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var ErrFileNotFound = errors.New("file not found")

func Extract(zipPath, filename string) (*Manifest, error) {
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

	return nil, ErrFileNotFound
}

func Hash(zipPath string) (string, error) {
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

type Index []IndexEntry

func ParseIndex(stripPrefix, path string) (Index, error) {
	var index []IndexEntry
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".zip") {
			manifest, err := Extract(path, ManifestFile)
			if err != nil {
				return err
			}

			if manifest != nil {
				hash, err := Hash(path)
				if err != nil {
					return err
				}

				relativePath, err := filepath.Rel(stripPrefix, path)
				if err != nil {
					return err
				}

				index = append(index, IndexEntry{Path: relativePath, Sha256: hash, Manifest: *manifest})
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return index, nil
}

func (index Index) Sort() {
	sort.Slice(index, func(i, j int) bool {
		// parse rfc3339 date from string
		iDate, err := time.Parse(time.RFC3339, index[i].Manifest.Created)
		if err != nil {
			return false
		}
		jDate, err := time.Parse(time.RFC3339, index[j].Manifest.Created)
		if err != nil {
			return false
		}
		return iDate.After(jDate)
	})
}

func (index Index) Write(stripPrefix, output string) error {
	for i := range index {
		index[i].Path = strings.TrimPrefix(index[i].Path, stripPrefix)
		index[i].Path = strings.TrimPrefix(index[i].Path, "/")
	}

	// keep it human readable
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
