package mapper

import (
	"bytes"
	"git.miem.hse.ru/1206/material-library/internal/dto"
	"git.miem.hse.ru/1206/proto/pb"
)

func UploadMaterialRequest(request *pb.UploadMaterialRequest, size int64, reader *bytes.Reader) dto.UploadMaterialRequest {
	var dto dto.UploadMaterialRequest
	dto.Author = request.GetInfo().Author
	dto.Title = request.GetInfo().Title
	dto.FileType = request.GetInfo().Type
	dto.FileLink = request.GetInfo().Link
	dto.Reader = reader
	return dto
}
