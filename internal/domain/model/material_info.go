package model

import "github.com/google/uuid"

type MaterialInfo struct {
	MaterialID uuid.UUID
	Author     string
	Title      string
	FileType   string
	FileLink   string
}
