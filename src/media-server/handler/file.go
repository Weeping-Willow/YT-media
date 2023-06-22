package handler

import (
	"net/http"

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
}

func NewFileHandler(fileService service.FileService) FileHandler {
	return fileHandler{
		fileService: fileService,
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
