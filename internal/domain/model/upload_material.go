package model

import "bytes"

type UploadMaterial struct {
	title string
	data  bytes.Buffer
}
