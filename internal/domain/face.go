package domain

import (
	"git.miem.hse.ru/1206/material-library/internal/domain/model"
	"git.miem.hse.ru/1206/material-library/internal/dto"
)

type StorageRepo interface {
	UploadMaterial(material *model.Material) (string, error)
	DeleteMaterialByObjectName(objectName string) error
	GetMaterialByObjectName(objectName string) (*model.Material, error)
}

type MaterialInfoRepo interface {
	UploadMaterialInfo(info *model.MaterialInfo) error
}

type Storage interface {
	UploadMaterial(material *dto.UploadMaterialRequest) (*dto.UploadMaterialResponse, error)
}
