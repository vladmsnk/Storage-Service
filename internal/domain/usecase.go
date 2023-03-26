package domain

import (
	"git.miem.hse.ru/1206/material-library/internal/dto"
	"github.com/google/uuid"
)

type UseCase struct {
	storageRepo  StorageRepo
	materialRepo MaterialInfoRepo
}

func NewUseCase(r StorageRepo, m MaterialInfoRepo) *UseCase {
	return &UseCase{storageRepo: r, materialRepo: m}
}

func (u *UseCase) UploadMaterial(m *dto.UploadMaterialRequest) (*dto.UploadMaterialResponse, error) {

	uploadMaterial := m.FromDTO(uuid.New())

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
