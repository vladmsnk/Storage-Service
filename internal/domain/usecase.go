package domain

type UseCase struct {
	storageRepo  StorageRepo
	materialRepo MaterialInfoRepo
}

func NewUseCase(r StorageRepo, m MaterialInfoRepo) *UseCase {
	return &UseCase{storageRepo: r, materialRepo: m}
}
