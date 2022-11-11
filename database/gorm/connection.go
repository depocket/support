package connection

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"moul.io/zapgorm2"
)

type Config struct {
	URL         string `env:"DATABASE_URL"`
	Host        string `env:"DATABASE_HOST"`
	Port        int32  `env:"DATABASE_PORT"`
	User        string `env:"DATABASE_USER"`
	Password    string `env:"DATABASE_PASSWORD"`
	Database    string `env:"DATABASE_DB"`
	Environment string `env:"GIN_MODE"`
	AppName     string `env:"APP_NAME"`
}

func NewConnection(log *zap.Logger, sch *string) *gorm.DB {
	config := parseConfig()
	connection := config.ToConnectionURI()
	var logger *zapgorm2.Logger = nil
	if log != nil {
		lg := zapgorm2.New(log)
		lg.SetAsDefault()
		logger = &lg
	}

	public := "public"
	if sch == nil {
		sch = &public
	}
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   *sch + ".",
			SingularTable: false,
		},
		Logger: logger,
	})
	if err != nil {
		log.Error(err.Error())
	}
	return db
}

func (c Config) ToConnectionURI() string {
	return fmt.Sprintf(`postgresql://%s:%s@%s:%v/%s?application_name=%s`, c.User, c.Password, c.Host, c.Port, c.Database, c.AppName)
}

func parseConfig() *Config {
	c := Config{}
	if err := env.Parse(&c); err != nil {
		log.Fatal(err)
	}
	return &c
}
