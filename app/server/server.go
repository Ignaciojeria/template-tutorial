package server

import (
	"archetype/app/configuration"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
)

var _ = ioc.Registry(NewServer, configuration.NewConf)

type Server struct {
	e    *echo.Echo
	conf configuration.Conf
}

func NewServer(conf configuration.Conf) Server {
	return Server{
		e:    echo.New(),
		conf: conf,
	}
}

func (s Server) Start() {
	s.e.Start(":" + s.conf.Port)
}

func (s Server) Router() *echo.Echo {
	return s.e
}
