package main

import (
	"os"
	"path/filepath"
)

type FileDB struct {
	path string
}

func NewFileDB(path string) *FileDB {
	return &FileDB{
		path: filepath.Dir(path),
	}
}

func (f *FileDB) getDirs(paths ...string) ([]string, error) {
	path := filepath.Join(f.path, filepath.Join(paths...))

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dirNames []string
	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			return nil, err
		}
		if fileInfo.IsDir() {
			dirNames = append(dirNames, fileInfo.Name())
		}
	}

	return dirNames, nil
}

func (f *FileDB) getFiles(paths ...string) ([]string, error) {
	path := filepath.Join(f.path, filepath.Join(paths...))

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dirNames []string
	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			return nil, err
		}
		if !fileInfo.IsDir() {
			dirNames = append(dirNames, fileInfo.Name())
		}
	}

	return dirNames, nil
}

func GetChains()
