package httputil

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

func BindStrict(ctx echo.Context, destination interface{}) map[string]string {
	if err := ctx.Bind(destination); err != nil {
		return ErrorMsgToMap(http.StatusText(http.StatusBadRequest))
	}
	if isValid, errs := Validate(destination); !isValid {
		return errs
	}
	return nil
}

func Validate(s interface{}) (bool, map[string]string) {
	ok, err := govalidator.ValidateStruct(s)
	if ok {
		return true, nil
	}

	errs := make(map[string]string)
	for _, item := range err.(govalidator.Errors).Errors() {
		n := govalidator.ErrorsByField(item)
		for k, v := range n {
			errs[k] = v
		}
	}
	return false, errs
}
