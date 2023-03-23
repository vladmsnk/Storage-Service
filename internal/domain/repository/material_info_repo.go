package repository

import (
	"git.miem.hse.ru/1206/app/storage/stpg"
	"git.miem.hse.ru/1206/material-library/internal/domain/model"
)

// MaterialInfoRepo repository with postgres connection
type MaterialInfoRepo struct {
	postgres stpg.PGer // postgres instance
}

func NewMaterialInfoRepo() *MaterialInfoRepo {
	return &MaterialInfoRepo{postgres: stpg.Gist()}
}

func (r *MaterialInfoRepo) UploadMaterialInfo(info *model.UploadMaterialInfo) error {
	return nil
}
