package domain

import (
	"git.miem.hse.ru/1206/material-library/internal/domain/model"
)

type StorageRepo interface {
	UploadMaterial(material *model.UploadMaterial) (int64, error)
}

type MaterialInfoRepo interface {
	UploadMaterialInfo(info *model.UploadMaterialInfo) error
}
