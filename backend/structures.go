package main

type FileIndex struct {
	hash          string
	file          []byte
	size          int
	expires       int
	auto_delete   bool
	max_downloads int
	downloads     int
	password      string
}

type PasteIndex struct {
	hash          string
	text          string
	language      string
	expires       int
	auto_delete   *bool
	max_downloads int
	downloads     int
	password      string
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type DeleteResponse struct {
	success bool
	message string
}
