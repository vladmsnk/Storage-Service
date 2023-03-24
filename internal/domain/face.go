package domain

import (
	"git.miem.hse.ru/1206/material-library/internal/domain/model"
	"git.miem.hse.ru/1206/material-library/internal/dto"
)

type StorageRepo interface {
	UploadMaterial(material *model.UploadMaterial) (string, error)
	DeleteMaterialByObjectName(objectName string) error
}

type MaterialInfoRepo interface {
	UploadMaterialInfo(info *model.UploadMaterialInfo) error
}

type Storage interface {
	UploadMaterial(material *dto.UploadMaterialRequest) (*dto.UploadMaterialResponse, error)
}
