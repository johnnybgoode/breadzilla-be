package database

import (
	"context"
	"log"

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

func (c *Config) ProcessFromEnv(ctx context.Context) {
	if err := envconfig.Process(ctx, c); err != nil {
		log.Fatal(err)
	}
}

func (c *Config) GetMysql() *mysql.Config {
	m := mysql.Config{User: c.User, Passwd: c.Passwd, Net: c.Net, Addr: c.Addr, DBName: c.DBName}
	return &m
}

// func (c *Config) GetMysqlConfig() {
// 	v := reflect.ValueOf(*c)
// 	fmt.Printf("GetMysqlConfig reflected %v", v)
// 	// fmt.Println(v.Field(0).Interface())
// 	cfg := mysql.NewConfig()
// 	cfg.Apply(func (conf *mysql.Config) error {
//     for i := 0; i < v.NumField(); i++ {
// 			n := v.Field(i)
// 			fmt.Println(n)
//       // conf.{v.Field(i)} = v.Field(i).Interface()
//     }
// 		return nil
// 	})
// }

// func (c *Config) ConnectionURL() string {
// 	if c == nil {
// 		return ""
// 	}
// 	host := c.Host
// 	if v := c.Port; v != "" {
// 		host = host + ":" + v
// 	}
// 	u := &url.URL{
// 		Scheme: "mysql",
// 		Host:   host,
// 		Path:   c.DBName,
// 	}
// 	if c.User != "" || c.Password != "" {
// 		u.User = url.UserPassword(c.User, c.Password)
// 	}
// 	return u.String()
// }
