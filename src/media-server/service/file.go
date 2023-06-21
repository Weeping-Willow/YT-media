package service

import (
	"context"

	"media-server/repository"
)

type FileService interface {
	Exist(c context.Context, path string) (bool, error)
	Update(c context.Context) error
	Create(c context.Context) error
	Delete(c context.Context, path string) error
	Get(c context.Context, path string) error
}

type fileService struct {
	fileRepository repository.FileRepository
}

func NewFileService(fileRepository repository.FileRepository) FileService {
	return &fileService{
		fileRepository: fileRepository,
	}
}

func (f fileService) Exist(c context.Context, path string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (f fileService) Update(c context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f fileService) Create(c context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f fileService) Delete(c context.Context, path string) error {
	//TODO implement me
	panic("implement me")
}

func (f fileService) Get(c context.Context, path string) error {
	//TODO implement me
	panic("implement me")
}
