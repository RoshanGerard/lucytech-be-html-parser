package models

type HtmlParseResponseDto struct {
	Title         string `json:"title"`
	Version       string `json:"version"`
	InternalLinks int    `json:"internalLinks"`
	ExternalLinks int    `json:"externalLinks"`
	IsLoginPage   bool   `json:"isLoginPage"`
}
