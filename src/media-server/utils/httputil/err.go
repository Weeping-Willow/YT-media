package httputil

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrLevel int

const UserErr = ErrLevel(400)
const UnauthorizedErr = ErrLevel(401)
const ForbiddenErr = ErrLevel(403)
const NotFoundErr = ErrLevel(404)
const ServerErr = ErrLevel(500)

type HTTPError struct {
	Err   error
	Level ErrLevel
}

type ErrorResponse struct {
	Message string `json:"error"`
}

func (b HTTPError) Error() string {
	if b.Err != nil {
		return b.Err.Error()
	}
	return ""
}

func HandleError(context echo.Context, err error) error {
	httpError, ok := err.(HTTPError)
	if ok {
		switch httpError.Level {
		case UserErr:
			return BadRequestWithError(context, httpError)
		case UnauthorizedErr:
			return UnauthorizedWithError(context, httpError)
		case ForbiddenErr:
			return ForbiddenWithError(context, httpError)
		case NotFoundErr:
			return NotFoundWithError(context, httpError)
		case ServerErr:
			return InternalServerErrorWithError(context, httpError)
		}
	}
	if !ok && err != nil {
		return context.JSON(http.StatusInternalServerError, ErrorMsgToMap("something went wrong"))
	}
	return err
}

func BadRequestWithError(context echo.Context, httpError HTTPError) error {
	return context.JSON(http.StatusBadRequest, ErrorResponse{httpError.Error()})
}

func UnauthorizedWithError(context echo.Context, httpError HTTPError) error {
	return context.JSON(http.StatusUnauthorized, ErrorResponse{httpError.Error()})
}

func ForbiddenWithError(context echo.Context, httpError HTTPError) error {
	return context.JSON(http.StatusForbidden, ErrorResponse{httpError.Error()})
}

func NotFoundWithError(context echo.Context, httpError HTTPError) error {
	return context.JSON(http.StatusNotFound, ErrorResponse{httpError.Error()})
}

func InternalServerErrorWithError(context echo.Context, httpError HTTPError) error {
	return context.JSON(http.StatusInternalServerError, ErrorResponse{httpError.Error()})
}

func BadRequestErrorWithMsg(message string) HTTPError {
	return HTTPError{Err: errors.New(message), Level: UserErr}
}

func UnauthorizedErrorWithMsg(message string) HTTPError {
	return HTTPError{Err: errors.New(message), Level: UnauthorizedErr}
}

func ForbiddenErrorWithMsg(message string) HTTPError {
	return HTTPError{Err: errors.New(message), Level: ForbiddenErr}
}

func NotFoundErrorWithMsg(message string) HTTPError {
	return HTTPError{Err: errors.New(message), Level: NotFoundErr}
}

func ServerErrorWithMsg(message string) HTTPError {
	return HTTPError{Err: errors.New(message), Level: ServerErr}
}

func ErrorMsgToMap(err string) map[string]string {
	return map[string]string{"error": err}
}
