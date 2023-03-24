package model

import "bytes"

type UploadMaterial struct {
	ObjectName  string
	ObjectSize  int64
	ContentType string
	Reader      *bytes.Reader
}
