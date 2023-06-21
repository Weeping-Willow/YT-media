package router

import "github.com/labstack/echo/v4"

const GroupFile = "file"

func (r *router) RegisterFileRoutes(g *echo.Group) {
	fileGroup := g.Group(GroupFile)

	fileGroup.GET("/exist/:path", r.fileHandler.Exist)
	fileGroup.POST("", r.fileHandler.Create)
	fileGroup.PATCH("", r.fileHandler.Update)
	fileGroup.DELETE("/exist/:path", r.fileHandler.Delete)
	fileGroup.GET("/exist/:path", r.fileHandler.Get)
}
