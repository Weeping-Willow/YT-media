package service

import (
	"context"

	"media-server/repository"
	"media-server/utils/ctxutil"
	"media-server/utils/httputil"
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

var (
	ErrCheckIfFileExists = httputil.ServerErrorWithMsg("check if file exists")
)

func NewFileService(fileRepository repository.FileRepository) FileService {
	return &fileService{
		fileRepository: fileRepository,
	}
}

func (f fileService) Exist(c context.Context, path string) (bool, error) {
	logger := ctxutil.GetLogger(c)
	exists, err := f.fileRepository.Exist(c, path)
	if err != nil {
		logger.WithError(err).Error("Failed to check file existence")
		return false, ErrCheckIfFileExists
	}

	return exists, nil
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
