package model

import "bytes"

type Material struct {
	ObjectName  string
	ObjectSize  int64
	ContentType string
	Size        int64
	Reader      *bytes.Reader
}
