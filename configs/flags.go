package configs

import "github.com/pborman/getopt/v2"

type Flags struct {
	ConfigPath *string
}

func ParseFlags() (fs *Flags) {
	fs = &Flags{
		ConfigPath: getopt.StringLong("config", 'c', "", "Specify custom config path"),
	}
	getopt.ParseV2()
	return
}
