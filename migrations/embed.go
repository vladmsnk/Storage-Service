package migrations

import (
	"embed"
	"path"
)

// Здесь обычно находятся методы, отвечающие за накатывание миграций

var (
	//go:embed postgres/*.up.sql
	UpMigrations embed.FS

	//go:embed postgres/*.down.sql
	DownMigrations embed.FS
)

func GetAllFilenames(fs *embed.FS, dir string) (out []string, err error) {
	if len(dir) == 0 {
		dir = "."
	}

	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fp := path.Join(dir, entry.Name())
		if entry.IsDir() {
			res, err := GetAllFilenames(fs, fp)
			if err != nil {
				return nil, err
			}

			out = append(out, res...)

			continue
		}

		out = append(out, fp)
	}

	return
}
