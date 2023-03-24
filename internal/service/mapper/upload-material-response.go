package mapper

import (
	"git.miem.hse.ru/1206/material-library/internal/dto"
	"git.miem.hse.ru/1206/proto/pb"
)

func UploadMaterialResponse(response dto.UploadMaterialResponse) *pb.UploadMaterialResponse {
	return &pb.UploadMaterialResponse{Id: response.MaterialID, Size: response.Size}
}
