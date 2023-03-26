package dto

import "github.com/google/uuid"

type UploadMaterialResponse struct {
	MaterialID uuid.UUID
	Size       int64
}
