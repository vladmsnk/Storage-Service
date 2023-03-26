package domain

import (
	"git.miem.hse.ru/1206/material-library/internal/domain/model"
	"git.miem.hse.ru/1206/material-library/internal/dto"
	"github.com/google/uuid"
)

type StorageRepo interface {
	UploadMaterial(material *model.Material) (uuid.UUID, error)
	DeleteMaterialByObjectName(objectName string) error
	GetMaterialByObjectName(objectName string) (*model.Material, error)
	CheckMaterialExists(objectName string) (bool, error)
}

type MaterialInfoRepo interface {
	UploadMaterialInfo(info *model.MaterialInfo) error
}

type Storage interface {
	UploadMaterial(material *dto.UploadMaterialRequest) (*dto.UploadMaterialResponse, error)
}
