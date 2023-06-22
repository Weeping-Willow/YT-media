package handler

import (
	"net/http"
	"path/filepath"

	"media-server/config"
	"media-server/service"
	"media-server/structures/dto"
	"media-server/utils/httputil"

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
	conf        config.Storage
}

func NewFileHandler(fileService service.FileService, conf config.Storage) FileHandler {
	return fileHandler{
		fileService: fileService,
		conf:        conf,
	}
}

func (f fileHandler) Exist(ctx echo.Context) error {
	var req dto.FileExitsRequest
	if errs := httputil.BindStrict(ctx, &req); errs != nil {
		return ctx.JSON(http.StatusBadRequest, errs)
	}

	exists, err := f.fileService.Exist(ctx.Request().Context(), req.Path)
	if err != nil {
		return httputil.HandleError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, dto.FileExitsResponse{
		Exists: exists,
	})

}

func (f fileHandler) Update(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f fileHandler) Create(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f fileHandler) Delete(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (f fileHandler) Get(ctx echo.Context) error {
	var req dto.FileExitsRequest
	if errs := httputil.BindStrict(ctx, &req); errs != nil {
		return ctx.JSON(http.StatusBadRequest, errs)
	}

	return ctx.File(f.conf.Path + filepath.Clean(req.Path))
}
