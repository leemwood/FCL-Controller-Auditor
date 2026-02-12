package models

type IndexEntry struct {
	ID           string   `json:"id"`
	Lang         string   `json:"lang"`
	Name         string   `json:"name"`
	Introduction string   `json:"introduction"`
	Device       []int    `json:"device"`
	Categories   []int    `json:"categories"`
}

type Category struct {
	ID   int `json:"id"`
	Lang []struct {
		Locale string `json:"locale"`
		Text   string `json:"text"`
	} `json:"lang"`
}

type RepoVersion struct {
	Screenshot  int      `json:"screenshot"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	Latest      Version  `json:"latest"`
	History     []Version `json:"history"`
}

type Version struct {
	VersionCode int    `json:"versionCode"`
	VersionName string `json:"versionName"`
}
