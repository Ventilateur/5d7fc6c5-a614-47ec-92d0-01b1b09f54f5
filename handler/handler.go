package handler

import (
	"github.com/Ventilateur/dataimpact-test/database"
	"github.com/Ventilateur/dataimpact-test/filesystem"
)

type Handler struct {
	db        *database.DB
	fs        *filesystem.FileStorage
	tokenizer *Tokenizer
}

func (h *Handler) Init(db *database.DB, fs *filesystem.FileStorage, tokenizer *Tokenizer) {
	h.db = db
	h.fs = fs
	h.tokenizer = tokenizer
}
