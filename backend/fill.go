package main

import (
	"time"
)

func FillPaste(paste *PasteRequest) PasteIndex {
	ret := new(PasteIndex)

	if paste.Language == "" {
		ret.Language = "text"
	} else {
		ret.Language = paste.Language
	}

	if paste.Expires_At == 0 {
		paste.Expires_At = time.Now().Add(time.Hour * 24 * 7 * 3).Unix()
	}

	ret.Text = paste.Text
	ret.Expires_At = paste.Expires_At
	ret.Password = paste.Password
	ret.Auto_Delete = paste.Auto_Delete
	ret.Max_Downloads = paste.Max_Downloads
	return *ret
}
