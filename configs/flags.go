package configs

import "flag"

type Flags struct {
	ConfigPath string
}

func parseFlags() (fs *Flags) {
	fs = new(Flags)
	flag.StringVar(&fs.ConfigPath, "config", "", "specify custom config path")
	flag.Parse()
	return
}
