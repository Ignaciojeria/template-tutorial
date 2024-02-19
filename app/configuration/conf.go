package configuration

import (
	"os"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/joho/godotenv"
)

var _ = ioc.Registry(NewConf)

type Conf struct {
	Port      string
	ApiPrefix string
}

func NewConf() (Conf, error) {
	err := godotenv.Load()
	if err != nil {
		return Conf{}, err
	}
	port := os.Getenv("Port")
	apiPrefix := os.Getenv("ApiPrefix")
	conf := Conf{
		Port:      port,
		ApiPrefix: apiPrefix,
	}
	return conf, nil
}
