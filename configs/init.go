package configs

import (
	_ "embed"
	"io"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/sharpvik/log-go/v2"
)

//go:embed default.yml
var defaultConfigFile []byte

type Config struct {
	Database *Database
	Server   *Server
}

type Database struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type Server struct {
	Address string
}

func MustInit() (config *Config) {
	log.Debug("reading config ...")
	defer log.Debug("config successfull")
	flags := ParseFlags()
	if *flags.ConfigPath == "" {
		return defaultConfig()
	}
	return customConfig(os.Open(*flags.ConfigPath))
}

func defaultConfig() (config *Config) {
	config = new(Config)
	if err := yaml.Unmarshal(defaultConfigFile, config); err != nil {
		log.Fatal("failed to read default config file: %s", err)
	}
	return
}

func customConfig(file io.Reader, err error) (config *Config) {
	config = new(Config)
	if err := yaml.NewDecoder(file).Decode(config); err != nil {
		log.Fatal("failed to read default config file: %s", err)
	}
	return
}
