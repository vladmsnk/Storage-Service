package domain

import (
	"git.miem.hse.ru/1206/material-library/internal/domain/model"
)

// Файлы face.go предназначены для размещения в них интерфейсов.
// Напоминаю, что интерфейсы всегда располагаются рядом с использованием, а не рядом с реализацией.
type StorageRepo interface {
	UploadMaterial(material *model.UploadMaterial) (int64, error)
}

type MaterialInfoRepo interface {
	UploadMaterialInfo(info *model.UploadMaterialInfo) error
}
