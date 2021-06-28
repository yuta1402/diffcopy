package diffcopy

import (
	"os"
	"path/filepath"
)

func FindWaitingFiles(srcDir string, destDir string) ([]string, error) {
	list := []string{}

	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		srcPath := path

		path, _ = filepath.Rel(srcDir, path)
		destPath := filepath.Join(destDir, path)

		// ファイルが存在している場合
		if _, err := os.Stat(destPath); err == nil {
			return nil
		}

		list = append(list, srcPath)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return list, nil
}
