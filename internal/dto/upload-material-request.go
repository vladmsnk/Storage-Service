package dto

import (
	"bytes"
	"git.miem.hse.ru/1206/material-library/internal/domain/model"
)

type UploadMaterialRequest struct {
	Author   string
	Title    string
	FileType string
	FileLink string
	Size     int64
	Reader   *bytes.Reader
}

func (u UploadMaterialRequest) FromDTOInfo(materialId string) model.MaterialInfo {
	return model.MaterialInfo{MaterialID: materialId,
		Author:   u.Author,
		Title:    u.Title,
		FileType: u.FileType,
		FileLink: u.FileLink}
}

func (u UploadMaterialRequest) FromDTO() model.Material {
	return model.Material{
		ObjectName:  u.Title,
		ObjectSize:  u.Size,
		ContentType: u.FileType,
		Reader:      u.Reader}
}
