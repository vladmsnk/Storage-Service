package domain

import (
	"git.miem.hse.ru/1206/app/errs"
	"git.miem.hse.ru/1206/material-library/internal/body"
	"git.miem.hse.ru/1206/material-library/internal/dto"
)

type UseCase struct {
	storageRepo  StorageRepo
	materialRepo MaterialInfoRepo
}

func NewUseCase(r StorageRepo, m MaterialInfoRepo) *UseCase {
	return &UseCase{storageRepo: r, materialRepo: m}
}

func (u *UseCase) UploadMaterial(m *dto.UploadMaterialRequest) (*dto.UploadMaterialResponse, error) {

	uploadMaterial := m.FromDTO()

	object, err := u.storageRepo.GetMaterialByObjectName(uploadMaterial.ObjectName)
	if err != nil {
		return nil, err
	}
	if object != nil {
		return nil, errs.New(err, body.ErrAlreadyExists)
	}
	materialID, err := u.storageRepo.UploadMaterial(&uploadMaterial)
	if err != nil {
		return nil, err
	}
	uploadMaterialInfo := m.FromDTOInfo(materialID)
	err = u.materialRepo.UploadMaterialInfo(&uploadMaterialInfo)
	if err != nil {
		u.storageRepo.DeleteMaterialByObjectName(m.Title) //compensate
		return nil, err
	}

	return &dto.UploadMaterialResponse{MaterialID: materialID, Size: m.Size}, nil
}
