package server

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/johnnybgoode/breadzilla/pkg/common"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	RouteMap     map[string]RouteHandler
	RouteHandler func(echo.Context, *sql.DB) error

	Server struct {
		address string
		db      *sql.DB
		echo    *echo.Echo
	}
)

func NewServer(address string, db *sql.DB) *Server {
	e := echo.New()
	e.HTTPErrorHandler = errorHandler
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	return &Server{
		address: address,
		db:      db,
		echo:    e,
	}
}

func (s *Server) ApplyRoutes(routes RouteMap) *Server {
	e := reflect.ValueOf(s.echo)
	for mp, handler := range routes {
		method, path := common.Must2(parseMethodAndPath(mp))

		args := make([]reflect.Value, 2)
		args[0] = reflect.ValueOf(path)
		args[1] = reflect.ValueOf(s.withDB(handler))
		e.MethodByName(method).Call(args)
	}
	return s
}

func (s *Server) withDB(handler RouteHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handler(c, s.db)
	}
}

func (s *Server) Start() {
	s.echo.Logger.Fatal(s.echo.Start(s.address))
}

func (s *Server) DB() *sql.DB {
	return s.db
}

func (s *Server) Echo() *echo.Echo {
	return s.echo
}

func parseMethodAndPath(v string) (method string, path string, err error) {
	mp := strings.Split(v, "::")
	if len(mp) != 2 {
		return "", "", fmt.Errorf("Unable to parse route '%v'", v)
	}
	return mp[0], mp[1], nil
}
