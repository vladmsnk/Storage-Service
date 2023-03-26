package mapper

import (
	"git.miem.hse.ru/1206/material-library/internal/dto"
	"git.miem.hse.ru/1206/proto/pb"
)

func DeleteMaterialRequest(r *pb.DeleteMaterialRequest) dto.DeleteMaterialRequest {
	return dto.DeleteMaterialRequest{MaterialID: r.MaterialId}
}
