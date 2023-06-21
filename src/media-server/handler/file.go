package handler

import (
	"media-server/service"

	"github.com/labstack/echo/v4"
)

type FileHandler interface {
	Exist(echo.Context) error
	Update(echo.Context) error
	Create(echo.Context) error
	Delete(echo.Context) error
	Get(echo.Context) error
}

type fileHandler struct {
	fileService service.FileService
}

func NewFileHandler(fileService service.FileService) FileHandler {
	return fileHandler{
		fileService: fileService,
	}
}

func (f fileHandler) Exist(context echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f fileHandler) Update(context echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f fileHandler) Create(context echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f fileHandler) Delete(context echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f fileHandler) Get(context echo.Context) error {
	//TODO implement me
	panic("implement me")
}
