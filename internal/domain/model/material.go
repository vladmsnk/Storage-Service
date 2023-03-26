package model

import (
	"bytes"
	"github.com/google/uuid"
)

type Material struct {
	ObjectName  uuid.UUID
	ObjectSize  int64
	ContentType string
	Size        int64
	Reader      *bytes.Reader
}
