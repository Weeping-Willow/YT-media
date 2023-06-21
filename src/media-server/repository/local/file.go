package local_storage

import (
	"context"

	"media-server/repository"
)

type fileRepository struct {
}

func NewFileRepository() repository.FileRepository {
	return &fileRepository{}
}

func (f fileRepository) Exist(c context.Context, path string) (bool, error) {
	//TODO implement me
	panic("implement me")
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
