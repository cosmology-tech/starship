package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

type FileDB struct {
	logger *zap.Logger
	path   string
}

func NewFileDB(logger *zap.Logger, path string) *FileDB {
	dirPath := filepath.Dir(path)
	if err := os.MkdirAll(dirPath, os.ModeDir); err != nil {
		panic(fmt.Sprintf("Unable to create root directory at: %s", dirPath))
	}
	logger.Info("Initialized directory used for storing files",
		zap.String("path", dirPath))

	return &FileDB{
		path:   dirPath,
		logger: logger,
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

	var fileNames []string
	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			return nil, err
		}
		if !fileInfo.IsDir() {
			fileNames = append(fileNames, fileInfo.Name())
		}
	}

	return fileNames, nil
}

// GetChains will return a list of chains reading from files
func (f *FileDB) GetChains() ([]Chain, error) {
	chainDirs, err := f.getDirs("/")
	if err != nil {
		return nil, err
	}

	var chains []Chain
	for _, chainDir := range chainDirs {
		validators, err := f.GetChainValidators(chainDir)
		if err != nil {
			return nil, err
		}
		chains = append(chains, Chain{
			Name:       chainDir,
			Validators: validators,
		})
	}

	return chains, nil
}

func (f *FileDB) IsChain(name string) bool {
	chains, err := f.GetChains()
	if err != nil {
		return false
	}

	for _, chain := range chains {
		if chain.Name == name {
			return true
		}
	}

	return false
}

func (f *FileDB) CreateChain(name string) error {
	err := os.Mkdir(filepath.Join(f.path, name), os.ModeDir)
	if os.IsExist(err) {
		return nil
	}

	return err
}

// GetChainValidators will return a list of validators in the chain
func (f *FileDB) GetChainValidators(chain string) ([]Validator, error) {
	valDirs, err := f.getDirs(chain)
	if err != nil {
		return nil, err
	}

	var vals []Validator
	for _, valDir := range valDirs {
		vals = append(vals, Validator{
			Name:    valDir,
			Moniker: valDir,
		})
	}

	return vals, nil
}

func (f *FileDB) IsValidator(chain string, name string) bool {
	vals, err := f.GetChainValidators(chain)
	if err != nil {
		return false
	}

	for _, val := range vals {
		if val.Name == name {
			return true
		}
	}

	return false
}

func (f *FileDB) CreateValidator(chain string, name string) error {
	err := os.Mkdir(filepath.Join(f.path, chain, name), os.ModeDir)
	if os.IsExist(err) {
		return nil
	}

	return err
}

func (f *FileDB) ListSnapshots(chain string, validator string) ([]string, error) {
	if !f.IsChain(chain) {
		return nil, fmt.Errorf("chain %s does not exists", chain)
	}
	if !f.IsValidator(chain, validator) {
		return nil, fmt.Errorf("validator %s does not exists for %s", chain, validator)
	}

	files, err := f.getFiles(chain, validator, "snapshots")
	if err != nil {
		return nil, err
	}

	return files, nil
}

// GetSnapshot return snapshot stored for the chain and validator
func (f *FileDB) GetSnapshot(chain string, validator string, snapshot string) ([]byte, error) {
	filePath := filepath.Join(f.path, chain, validator, "snapshots", snapshot)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (f *FileDB) StoreSnapshot(chain string, validator string, snapshot string, data []byte) error {
	err := os.MkdirAll(filepath.Join(f.path, chain, validator, "snapshots"), os.ModeDir)
	if err != nil {
		return err
	}

	filePath := filepath.Join(f.path, chain, validator, "snapshots", snapshot)
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("unable to write snapshot to %s, with err: %s", filePath, err)
	}

	return nil
}

func (f *FileDB) ListExports(chain string, validator string) ([]string, error) {
	if !f.IsChain(chain) {
		return nil, fmt.Errorf("chain %s does not exists", chain)
	}
	if !f.IsValidator(chain, validator) {
		return nil, fmt.Errorf("validator %s does not exists for %s", chain, validator)
	}

	files, err := f.getFiles(chain, validator, "exports")
	if err != nil {
		return nil, err
	}

	return files, nil
}

// GetExport return snapshot stored for the chain and validator
func (f *FileDB) GetExport(chain string, validator string, export string) ([]byte, error) {
	filePath := filepath.Join(f.path, chain, validator, "exports", export)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (f *FileDB) StoreExport(chain string, validator string, export string, data []byte) error {
	err := os.MkdirAll(filepath.Join(f.path, chain, validator, "exports"), os.ModeDir)
	if err != nil {
		return err
	}

	filePath := filepath.Join(f.path, chain, validator, "exports", export)
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("unable to write snapshot to %s, with err: %s", filePath, err)
	}

	return nil
}
