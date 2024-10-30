package sakescript

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	DefaultStoryType = "story"

	IndexFile    = "index.json"
	MainFile     = "main.json"
	ManifestFile = "manifest.json"
)

// Sentence struct
type Sentence struct {
	Ja string `json:"ja"`
	En string `json:"en"`
}

// MediaEntry struct
type MediaEntry struct {
	Type string `json:"type"`
	URI  string `json:"uri"`
}

// Chapter struct
type Chapter struct {
	Title        *Sentence          `json:"title,omitempty"`
	Sentences    []Sentence         `json:"sentences"`
	MediaEntries map[int]MediaEntry `json:"mediaEntries,omitempty"`
}

// Story struct
type Story struct {
	Title    Sentence    `json:"title"`
	Type     string      `json:"type,omitempty"`
	Cover    *MediaEntry `json:"cover,omitempty"`
	Chapters []Chapter   `json:"chapters"`
}

// validateStory checks the required fields and structures of the Story.
func (s *Story) Validate() error {
	if s.Title.En == "" || s.Title.Ja == "" {
		return fmt.Errorf("story title is incomplete, have en: %s, ja: %s", s.Title.En, s.Title.Ja)
	}

	storyType := DefaultStoryType
	if s.Type != "" {
		storyType = s.Type
	}
	if len(s.Chapters) == 0 {
		return errors.New("story has no chapters")
	}
	for _, c := range s.Chapters {
		if storyType == "story" && (c.Title == nil || (c.Title.En == "" || c.Title.Ja == "")) {
			return errors.New("chapter title is incomplete for story")
		}
		if len(c.Sentences) == 0 {
			return errors.New("chapter has no sentences")
		}
		for _, s := range c.Sentences {
			if s.En == "" || s.Ja == "" {
				return fmt.Errorf("sentence is incomplete, have en: %s, ja: %s", s.En, s.Ja)
			}
		}
	}

	return nil
}

// ParseStoryJSON takes a JSON string and attempts to parse it into a Story struct.
func (s *Story) FromString(jsonString string) error {
	if err := json.Unmarshal([]byte(jsonString), s); err != nil {
		return err
	}

	return nil
}

// Manifest struct
type Manifest struct {
	ID            string   `json:"id"`
	Version       string   `json:"version"`
	Type          string   `json:"type,omitempty"`
	Title         Sentence `json:"title"`
	TeaserImage   string   `json:"teaserImage,omitempty"`
	Author        string   `json:"author"`
	AuthorTwitter string   `json:"authorTwitter,omitempty"`
	AuthorNote    string   `json:"authorNote,omitempty"`
	Created       string   `json:"created"`
	Modified      string   `json:"modified"`
	Summary       Sentence `json:"summary"`
	Tags          []string `json:"tags"`
	License       string   `json:"license,omitempty"`
	Origin        string   `json:"origin,omitempty"`
}

// ValidateManifest checks the required fields and structures of the Manifest.
func (m *Manifest) Validate() error {
	if m.ID == "" {
		return errors.New("manifest ID is missing")
	}
	if m.Version == "" {
		return errors.New("manifest version is missing")
	}
	if m.Title.En == "" || m.Title.Ja == "" {
		return fmt.Errorf("manifest title is incomplete, have en: %s, ja: %s", m.Title.En, m.Title.Ja)
	}
	if m.Author == "" {
		return errors.New("manifest author is missing")
	}

	storyType := DefaultStoryType
	if m.Type != "" {
		storyType = m.Type
	}

	m.Created = time.Now().Format(time.RFC3339)
	m.Modified = time.Now().Format(time.RFC3339)

	if m.Summary.En == "" || m.Summary.Ja == "" {
		return errors.New("manifest summary is incomplete")
	}
	if len(m.Tags) == 0 {
		return errors.New("manifest has no tags")
	}

	if storyType == "article" {
		if m.Origin == "" {
			return errors.New("article has no origin")
		}
	}

	return nil
}

// FromString unmarshals a JSON string into a Manifest struct and validates it.
func (m *Manifest) FromString(jsonString string) error {
	if err := json.Unmarshal([]byte(jsonString), m); err != nil {
		return err
	}

	return nil
}

// IndexEntry struct
type IndexEntry struct {
	Path     string   `json:"path"`
	Sha256   string   `json:"sha256"`
	Manifest Manifest `json:"manifest"`
}
