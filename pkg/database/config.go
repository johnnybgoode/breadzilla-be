package database

import (
	"context"

	"github.com/go-sql-driver/mysql"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	User   string `env:"DB_USER" json:",omitempty"`
	Passwd string `env:"DB_PASS" json:"-"` // ignored by zap's JSON formatter
	Net    string `env:"DB_NET,  default=tcp" json:",omitempty"`
	Addr   string `env:"DB_HOST, default=127.0.0.1:3306" json:",omitempty"`
	DBName string `env:"DB_NAME, default=breadzilla" json:",omitempty"`
}

func (c *Config) DatabaseBConfig() *Config {
	return c
}

func (c *Config) ProcessFromEnv(ctx *context.Context) error {
	return envconfig.Process(*ctx, c)
}

func (c *Config) GetMysql() *mysql.Config {
	m := mysql.Config{User: c.User, Passwd: c.Passwd, Net: c.Net, Addr: c.Addr, DBName: c.DBName}
	return &m
}
