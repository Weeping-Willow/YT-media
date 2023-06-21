package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	conf "media-server/config"
	"media-server/handler"
	"media-server/repository"
	local_storage "media-server/repository/local"
	http_router "media-server/router"
	"media-server/service"

	"github.com/gookit/config"
	"github.com/gookit/config/yaml"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type App interface {
	Start() error
	Stop()
}

type app struct {
	config conf.Config

	fileService service.FileService

	fileHandler handler.FileHandler

	fileRepository repository.FileRepository
}

func NewApp() App {
	return &app{}
}

func (a *app) Start() error {
	if err := a.setup(); err != nil {
		return errors.Wrap(err, "failed to start app")
	}

	return nil
}

func (a *app) Stop() {
	fmt.Println("stopping app")
}

func (a *app) setup() error {
	logrus.SetOutput(os.Stdout)

	cfg, err := a.getConfig()
	if err != nil {
		return errors.Wrap(err, "failed to get config")
	}
	a.config = cfg

	a.fillRepository()
	a.fillService()
	a.fillHandlers()

	if err := a.startHTTPServer(); err != nil {
		return errors.Wrap(err, "failed to start http server")
	}

	return nil
}

func (a *app) getConfig() (conf.Config, error) {
	var cfg conf.Config

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	if err := config.LoadFiles("config.yaml"); err != nil {
		return conf.Config{}, errors.Wrap(err, "failed to load config")
	}

	cfg = conf.Config{
		App: conf.App{
			Name: "Media Server",
		},
		Http: conf.Http{
			Port: ":7000",
		},
		Storage: conf.Storage{
			Local: conf.Local{
				Path: "./local_storage/",
			},
		},
	}

	// TOOD: fix actually returing config
	return cfg, nil
}

func (a *app) fillRepository() {
	a.fileRepository = local_storage.NewFileRepository(a.config.Storage)
}

func (a *app) fillService() {
	a.fileService = service.NewFileService(a.fileRepository)
}

func (a *app) fillHandlers() {
	a.fileHandler = handler.NewFileHandler(a.fileService)
}

func (a *app) startHTTPServer() error {
	c := a.getEcho()

	router := http_router.NewRouter(a.fileHandler)
	router.RegisterRoutes(c)

	err := a.StartEcho(c)
	if err != nil {
		return errors.Wrap(err, "failed to start echo")
	}

	return nil
}

func (a *app) getEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodOptions, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	return e
}

func (a *app) StartEcho(e *echo.Echo) error {
	logrus.Infof("Starting HTTP server on: %s", a.config.Http.Port)
	listenConfig := net.ListenConfig{KeepAlive: a.config.Http.KeepAliveTimeout}
	listener, err := listenConfig.Listen(context.Background(), "tcp", a.config.Http.Port)
	if err != nil {
		return errors.Wrapf(err, "failed to setup TCP listener at %s", a.config.Port)
	}

	e.Listener = listener

	go func() {
		if err := e.Start(a.config.Port); err != nil && err != http.ErrServerClosed {
			logrus.WithError(err).Panic("Shutting down HTTP server")
		}
	}()

	return nil
}
