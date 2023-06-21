package local_storage

import (
	"context"
	"os"

	"media-server/config"
	"media-server/repository"
)

type fileRepository struct {
	conf config.Storage
}

func NewFileRepository(conf config.Storage) repository.FileRepository {
	return &fileRepository{
		conf: conf,
	}
}

func (f fileRepository) Exist(c context.Context, path string) (bool, error) {
	fullPath := f.getFilePath(path)
	_, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (f fileRepository) Update(c context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f fileRepository) Create(c context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f fileRepository) Delete(c context.Context, path string) error {
	//TODO implement me
	panic("implement me")
}

func (f fileRepository) Get(c context.Context, path string) error {
	//TODO implement me
	panic("implement me")
}

func (f fileRepository) getFilePath(path string) string {
	return f.conf.Path + path
}
