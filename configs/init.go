package configs

import (
	_ "embed"
	"encoding/json"
	"io"
	"os"

	"github.com/sharpvik/log-go/v2"
)

//go:embed default.json
var defaultConfigFile []byte

// Config contains configuration information for the whole service.
type Config struct {
	Database *Database `json:"database"`
	Server   *Server   `json:"server"`
}

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Server struct {
	Address string `json:"address"`
}

// MustInit attempts to initialise Config and panics in case of failure.
func MustInit() (config *Config) {
	log.Debug("reading config ...")
	defer log.Debug("config successfull")
	flags := parseFlags()
	if flags.ConfigPath == "" {
		return defaultConfig()
	}
	return customConfig(os.Open(flags.ConfigPath))
}

func defaultConfig() (config *Config) {
	config = new(Config)
	if err := json.Unmarshal(defaultConfigFile, config); err != nil {
		log.Fatal("failed to read default config file: %s", err)
	}
	return
}

func customConfig(file io.Reader, err error) (config *Config) {
	config = new(Config)
	if err := json.NewDecoder(file).Decode(config); err != nil {
		log.Fatal("failed to read default config file: %s", err)
	}
	return
}
