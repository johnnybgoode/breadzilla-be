package server

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	HandlerWithDBFunc func(echo.Context, *sql.DB) error
	RouteFunc         func(s *Server) error

	Server struct {
		address string
		db      *sql.DB
		echo    *echo.Echo
	}
)

func NewServer(address string, db *sql.DB) *Server {
	e := echo.New();
	e.HTTPErrorHandler = errorHandler;
	e.Use(middleware.RequestLogger())	
	e.Use(middleware.Recover())

	return &Server{
		address: address,
		db:      db,
		echo:    e,
	}
}

func (s *Server) ApplyRoutes(routeFuncs []RouteFunc) *Server {
	for i := range routeFuncs {
		routeFuncs[i](s)
	}
	return s
}

func (s *Server) Start() {
	s.echo.Logger.Fatal(s.echo.Start(s.address))
}

func (s *Server) WithDB(handlerFunc HandlerWithDBFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handlerFunc(c, s.db)
	}
}

func (s *Server) DB() *sql.DB {
	return s.db
}

func (s *Server) Echo() *echo.Echo {
	return s.echo
}