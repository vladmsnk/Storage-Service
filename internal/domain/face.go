package domain

// Файлы face.go предназначены для размещения в них интерфейсов.
// Напоминаю, что интерфейсы всегда располагаются рядом с использованием, а не рядом с реализацией.

type FileStorage interface {
	CreateFile(fileName string, fileData []byte) (id int64, err error)
	GetFile(id int64) (fileData []byte, err error)
	DeleteFile(id int64) error
}
