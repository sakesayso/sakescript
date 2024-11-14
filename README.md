# SakeScript Format Specification

[![Go Report Card](https://goreportcard.com/badge/github.com/sakesayso/sakescript)](https://goreportcard.com/report/github.com/sakesayso/sakescript)
[![GoDoc](https://godoc.org/github.com/sakesayso/sakescript?status.svg)](https://godoc.org/github.com/sakesayso/sakescript)

This document details the **SakeScript format**, a structured file format designed for the [SakeSaySo language learning app](https://sakesayso.com). SakeScript facilitates the packaging and distribution of various learning materials, such as stories and articles in a consistent and user-friendly manner. A SakeScript package is a ZIP archive containing a manifest file alongside the learning content and assets.

## SakeScript ZIP Archive Structure

Each SakeScript ZIP archive represents a single unit of learning content (e.g., a story, a news article, a lesson, or exercise). The archive includes:

- Manifest File: A `manifest.json` file containing metadata about the learning content.
- Content Files: `main.json` for stories and articles or `main.txt` for decks. Optionally various files (text, images, audio) constituting the learning material.

To create a SakeScript archive, you can use the zip cli. Ensure to include all necessary files (JSON files, images, etc.) in the archive. For example:
```sh
zip my-story-name.zip manifest.json main.json images/*
```

## main.json: Content File Format

The `main.json` file contains the main content of the learning material. SakeScript currently supports two types, 'story' and 'article'. The format for each is described below.

- title: A map of language codes to titles (e.g., "en": "The Mountain Trail").
- cover: This field supports image files. The `uri` can be a URL pointing to an external image (e.g., "https://example.org/cover.jpg") or a relative path to an image file within the archive (e.g., "images/cover.jpg"). For example:
```json
    "cover": {
        "type": "image",
        "uri": "images/cover.jpg" // or "https://example.org/cover.jpg"
    }
```
- type: Type of content, "story", "article" or "deck".
- chapters: List of chapters.
    + title (optional): Currently supported for 'story' type. A map of language codes to titles (e.g., "en": "About Tokyo").
    + sentences: List of sentences.
        * ja: Japanese sentence.
        * en: English sentence.

```json
{
    "title": {
        "en": "Journey Through Japan",
        "ja": "日本の旅"
    },
    "cover": {
        "type": "image",
        "uri": "https://www3.nhk.or.jp/news/html/20231111/K10014254991_2311111600_1111160953_01_02.jpg"
    },
    "type": "story",
    "chapters": [
        {
            "title": {
                "en": "About Tokyo",
                "ja": "東京について"
            },
            "sentences": [
                {
                    "ja": "東京は日本の首都です。",
                    "en": "Tokyo is the capital of Japan."
                },
                {
                    "ja": "新宿はにぎやかな場所です。",
                    "en": "Shinjuku is a bustling area."
                }
            ]
        }
    ]
}
```

## manifest.json: Metadata File Format

The `manifest.json` file in each SakeScript archive contains these fields:

- id: Unique script identifier for the content (e.g., UUID).
- type: Type of content, "story", "article" or "deck".
- version: Format version (e.g., "1.0").
- title: A map of language codes to titles (e.g., "en": "The Mountain Trail").
- created: Creation date, RFC3339 format (2020-12-29T12:00:00Z).
- modified: Last modification date, RFC3339 format (2020-12-29T12:00:00Z).
- author: Content author or creator.
- language: Primary language of the content.
- summary:  A map of language codes to summaries (e.g., "en": "A beginner-level story about a hike in the mountains.").
- license: License for the content (e.g., "Creative Commons").
- tags: List of tags for the content.
Optional fields:
- teaserImage (optional): Teaser image for the content.
- authorTwitter (optional): X/Twitter handle for the author.
- authorNote (optional): Author's note about the content.
- origin (optional): Source URL for the content.

Example
```json
{
    "id": "474007F8-F307-42F5-BA0E-E8B4547C7DAF",
    "type": "story",
    "version": "1.0",
    "title": {
        "en": "The Mountain Trail",
        "ja": "山道"
    },
    "author": "SakeSaySo",
    "authorTwitter": "sakesayso",
    "authorNote": "demo story",
    "teaserImage": "https://raw.githubusercontent.com/sakesayso/community/master/non-fiction/sci/2F98A92E-B14F-435F-B62E-2AD91FD0E862/cover.jpg",
    "created": "2020-12-13",
    "modified": "2023-12-13",
    "summary": {
        "en": "A beginner-level story about a hike in the mountains.",
        "ja": "初級者向けの山登りの話。"
    },
    "tags": [
        "BIZ",
        "N3"
    ],
    "license": "Creative Commons Attribution-ShareAlike",
    "origin": "https://www3.nhk.or.jp/news/easy/k10014288051000/k10014288051000.html"
}
```

Note: We recommend to use `uuidgen` or https://www.uuidgenerator.net/ or similar to generate an actually unique UUID.

If you include a cover image, we recommend using JPEG format for cover images to minimize file size. To convert a PNG image from e.g. DALL·E to JPEG, you can use ImageMagick with the following command:  `convert cover.png -resize 1080x -quality 92 cover.jpg`.

### Recommended Content Tags

Alongside JLPT levels (N1-N5), SakeScript supports arbitrary tags to categorize content. We recommend to use one JLPT level tag and at least one content tag.

Non-fiction content should use the following tags:
- **AME** - for arts, media, entertainment
- **TEC** - for technology, internet
- **SCI** - for science, environment
- **MED** - for health, medical, fitness
- **SPO** - for sports, esports
- **LIF** - for lifestyle, leasure
- **POL** - for politics, society
- **BIZ** - for finance, business, economics, military

Fiction content should use the following tags:
- **ADV** - for adventure, exploration
- **COM** - for comedy, humor
- **DRA** - for drama, relationships
- **DYS** - for dystopia, social Commentary
- **FAN** - for fantasy, mythology
- **HIS** - for historical, period
- **SFI** - for science fiction, futurism
- **THR** - for thriller, mystery

## Learning Decks for Space Repetition

A "deck" in SakeScript format contains a series of flash cards. Vocabulary can be matched in the dictionary upon import to gain cross-reference and translation benefits. In the manifest, a deck is denoted with the manifest type "deck". Instead of a json file, we then expect a `main.txt` that contains the deck in a simple text format. For example:

```
乾杯（かんぱい）
cheers, bottoms-up, prosit

宴会（えんかい）
party, banquet, reception
```

## Repository Index File

An `index.json` file is maintained in the repository to catalog all available SakeScript materials. This index, auto-generated from each archive's manifest file, includes:

- path: Relative path to the SakeScript ZIP in the repository.
- sha256: SHA-256 integrity hash of the ZIP archive.
- manifest: Extracted manifest data.

Example
```json
[
  {
    "path": "the-mountain-trail.zip",
    "sha256": "bf35415b1ee00fe56e6a8016848d7c7c35e392ca4732716dfce190a403b8303a",
    "manifest": {
      "id": "474007F8-F307-42F5-BA0E-E8B4547C7DAF",
      "version": "1.0",
      "title": {
        "en": "The Mountain Trail",
        "ja": "山道"
      },
      "author": "SakeSaySo",
      "authorTwitter": "sakesayso",
      "authorNote": "demo story",
      "created": "2020-12-13",
      "modified": "2023-12-13",
      "difficulty": "beginner",
      "summary": {
        "en": "A beginner-level story about a hike in the mountains.",
        "ja": "初級者向けの山登りの話。"
      },
      "tags": [
        "LIF",
        "N4"
      ],
      "license": "Creative Commons Attribution-ShareAlike"
    }
  }
  // ...
]
```

## Contribution and Usage Guidelines

See the content repository for more information on contributing and using SakeScript materials at https://github.com/sakesayso/community.

### Contributing to SakeScript

- Prepare your content and package it in a SakeScript ZIP file.
- Include a manifest.json file with accurate metadata.
- Place the ZIP file in the appropriate directory within the repository.
- Ensure the index.json is updated post-merge (typically automated).

### Content Licensing

We encourage the use of the "Creative Commons Attribution-ShareAlike" license. This license allows for both commercial and non-commercial use, modification, and distribution of content, as long as the original author is credited and any derivative works are shared under the same terms. This promotes a collaborative and open learning environment while ensuring creators receive recognition for their work.

#### How to License Your Content?

Simply include the "Creative Commons Attribution-ShareAlike" license in your manifest.json file. For more details on how to apply this license, visit [Creative Commons](https://creativecommons.org/).
