# SakeScript Format Specification

This document details the **SakeScript format**, a structured file format designed for the [SakeSaySo language learning app](https://sakesayso.com). SakeScript facilitates the packaging and distribution of various learning materials, such as stories and articles in a consistent and user-friendly manner. A SakeScript package is a ZIP archive containing a manifest file alongside the learning content and assets.

## SakeScript ZIP Archive Structure

Each SakeScript ZIP archive represents a single unit of learning content (e.g., a story, a news article, a lesson, or exercise). The archive includes:

- Manifest File: A `manifest.json` file containing metadata about the learning content.
- Content Files: `main.json` and various files (text, images, audio) constituting the learning material.

The zip cli can be used to create a SakeScript archive, e.g.:
```sh
zip my-story-name.zip *.json
```

## Manifest File Format

The `manifest.json` file in each SakeScript archive contains these fields:

- id: Unique script identifier for the content (e.g., UUID).
- type: Type of content (e.g., "story", "article").
- version: Format version (e.g., "1.0").
- title: A map of language codes to titles (e.g., "en": "The Mountain Trail").
- created: Creation date (YYYY-MM-DD).
- modified: Last modification date (YYYY-MM-DD).
- author: Content author or creator.
- language: Primary language of the content.
- summary:  A map of language codes to summaries (e.g., "en": "A beginner-level story about a hike in the mountains.").
- license: License for the content (e.g., "Creative Commons").
- tags: List of tags for the content.
Optional fields:
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

Note: use `uuidgen` or similar to generate an actually unique UUID.

### Tags

Alongside JLPT levels, SakeScript supports a set of tags to categorize content. It is recommended to use at least one JLPT level tag and one content tag.

Non-fiction content should use the following tags:
- **AME** - for arts, media, entertainment
- **TEC** - for technology, internet
- **SCI** - for science, environment
- **MED** - for health, medical
- **SPO** - for sports
- **LIF** - for lifestyle, leasure
- **POL** - for politics, society
- **BIZ** - for finance, business, economics

Fiction content should use the following tags:
- **ADV** - for adventure, exploration
- **COM** - for comedy, humor
- **DRA** - for drama, relationships
- **DYS** - for dystopia, social Commentary
- **FAN** - for fantasy, mythology
- **HIS** - for historical, period
- **SFI** - for science fiction, futurism
- **THR** - for thriller, mystery

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

## main.json File Format

The `main.json` file contains the main content of the learning material. SakeScript currently supports two types, 'story' and 'article'. The format for each is described below.

- title: A map of language codes to titles (e.g., "en": "The Mountain Trail").
- cover: Currently only supports images. uri is either a relative path to an image file in the archive or a URL to an external image.
- type: Type of content (e.g., "story", "article").
- chapters: List of chapters.
    + title (optional): Currently only supported for 'story' type. A map of language codes to titles (e.g., "en": "About Tokyo").
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
