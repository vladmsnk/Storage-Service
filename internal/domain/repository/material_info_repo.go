package repository

import (
	"context"
	"git.miem.hse.ru/1206/app/storage/stpg"
	"git.miem.hse.ru/1206/material-library/internal/body"
	"git.miem.hse.ru/1206/material-library/internal/domain/model"
	"github.com/jmoiron/sqlx"
)

// MaterialInfoRepo repository with postgres connection
type MaterialInfoRepo struct {
	postgres stpg.PGer // postgres instance
}

func NewMaterialInfoRepo() *MaterialInfoRepo {
	return &MaterialInfoRepo{postgres: stpg.Gist()}
}

func (r *MaterialInfoRepo) UploadMaterialInfo(i *model.UploadMaterialInfo) error {
	err := r.postgres.QueryTx(context.Background(), func(tx *sqlx.Tx) error {
		_, err := tx.Exec(body.SQLInsertMaterialInfo, i.MaterialID, i.Author, i.Title, i.FileType, i.FileLink)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
