package main

/*
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

type DeleteResponse struct {
	success bool
	message string
}
*/

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type PasteResponse struct {
	Admin_Key    string `json:"admin_key"`
	Download_Url string `json:"download_url"`
}

// The response body for the POST /api/upload/paste
type PasteIndex struct {
	PasteRequest
	Id        string `json:"id"` // Come back to this
	Downloads int    `json:"downloads"`
}

// The request body for the POST /api/upload/paste endpoint
type PasteRequest struct {
	Text          string `json:"text" validate:"required"` // Required
	Language      string `json:"language"`                 // Default: "text"
	Expires_At    int64  `json:"expires_at"`               // Default: 3 days from upload
	Max_Downloads int    `json:"max_downloads"`            // Default: 0 (unlimited)
	Auto_Delete   bool   `json:"auto_delete"`              // Default: false
	Password      string `json:"password"`                 // No password: ""
}

// The toml library parses integers as floats, so we need this temporary type.
type PreConfig struct {
	Save_Dir    string
	Max_Storage float64
	Expiry      float64
	Db_User     string
	Db_Pass     string
	Db_Name     string
}

type Config struct {
	Save_Dir    string
	Max_Storage int64
	Expiry      int64
	Db_User     string
	Db_Pass     string
	Db_Name     string
}
