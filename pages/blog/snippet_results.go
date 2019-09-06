package blog

import "github.com/mhconradt/personal_website/pages"

type ArticleSnippet struct {
	ID        int    `json:"id,omitempty"`
	Timestamp pages.Timestamp  `json:"timestamp,omitempty"`
	Title     string `json:"title,omitempty"`
	Body      string `json:"body,omitempty"`
}

type Cursor struct {
	Forward int64 `json:"forward"`
	Reverse int64 `json:"reverse"`
	Count int64 `json:"count"`
	Term string `json:"term"`
	Index string
}

type SearchResults struct {
	Results []ArticleSnippet `json:"results"`
	Cursor `json:"cursor"`
}
