package router

import "github.com/labstack/echo/v4"

const APIPath = "api/v1"

type Router interface {
	RegisterRoutes(e *echo.Echo)
}

type router struct{}

func NewRouter() Router {
	return &router{}
}

func (r router) RegisterRoutes(e *echo.Echo) {
	g := e.Group(APIPath)
	g.GET("/health", func(c echo.Context) error {
		return c.String(200, "ok")
	})
}
