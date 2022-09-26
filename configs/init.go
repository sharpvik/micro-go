package configs

import (
	_ "embed"
	"log"
	"os"

	"gopkg.in/yaml.v3"
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
	log.Print("reading config ...")
	flags := ParseFlags()
	if *flags.ConfigPath == "" {
		return defaultConfig()
	}
	config = customConfig(*flags.ConfigPath)
	log.Print("config successful")
	return
}

func defaultConfig() (config *Config) {
	config = new(Config)
	if err := yaml.Unmarshal(defaultConfigFile, config); err != nil {
		log.Fatalln("failed to read default config file:", err)
	}
	return
}

func customConfig(name string) (config *Config) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln("failed to open default config file:", err)
	}
	config = new(Config)
	if err := yaml.NewDecoder(file).Decode(config); err != nil {
		log.Fatalln("failed to read default config file:", err)
	}
	return
}
